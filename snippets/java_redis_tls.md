
# Maven

```
<dependency>
    <groupId>redis.clients</groupId>
    <artifactId>jedis</artifactId>
    <version>3.6.0</version>
</dependency>
```

# Java 1

```
package com.github.anigkus.redistls;

import java.io.FileInputStream;
import java.io.InputStream;
import java.security.KeyStore;
import java.security.SecureRandom;

import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSocketFactory;
import javax.net.ssl.TrustManager;
import javax.net.ssl.TrustManagerFactory;

import org.apache.commons.pool2.impl.GenericObjectPoolConfig;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;

public class JedisSSLTest {
    private static SSLSocketFactory createTrustStoreSSLSocketFactory(String jksFile) throws Exception {
        KeyStore trustStore = KeyStore.getInstance("jks");
        InputStream inputStream = null;
        try {
            inputStream = new FileInputStream(jksFile);
            trustStore.load(inputStream, "654321".toCharArray());//password
        } finally {
            inputStream.close();
        }
        
        TrustManagerFactory trustManagerFactory = TrustManagerFactory.getInstance("PKIX");
        trustManagerFactory.init(trustStore);
        
        TrustManager[] trustManagers = trustManagerFactory.getTrustManagers();

        SSLContext sslContext = SSLContext.getInstance("TLS");
        sslContext.init(null, trustManagers, new SecureRandom());
        return sslContext.getSocketFactory();
    }

    public static void main(String[] args) throws Exception {
        final SSLSocketFactory sslSocketFactory = createTrustStoreSSLSocketFactory("redis.jks");
        JedisPool pool = new JedisPool(new GenericObjectPoolConfig(), "x.x.x.x", 6379, 20000, null, 0, true, sslSocketFactory,null, null);

        try (Jedis jedis = pool.getResource()) {
            jedis.set("key", "value");
            System.out.println(jedis.get("key"));
        }
    }
}
```

# Java 2

```
package com.github.anigkus.redistls;

import redis.clients.jedis.Jedis;

public class RedisDemo {
    public static void main(String[] args) {
        setupTrustStore();
        Jedis jedis = new Jedis("rediss://x.x.x.x:6379");
        String pong = jedis.ping();
        System.out.println("ping result:" + pong);
        jedis.close();
    }

    private static void setJvmTrustStore(String trustStoreFilePath, String trustStoreType) {
        System.setProperty("javax.net.ssl.trustStore", trustStoreFilePath);
        System.setProperty("javax.net.ssl.trustStoreType", trustStoreType);
        System.setProperty("javax.net.ssl.trustStorePassword", "654321");
    }


    private static void setupTrustStore() {
        setJvmTrustStore("redis.jks", "jks");
    }
}

```

# Errors

## E1

```
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
Exception in thread "main" redis.clients.jedis.exceptions.JedisConnectionException: javax.net.ssl.SSLHandshakeException: Received fatal alert: certificate_required
        at redis.clients.jedis.util.RedisInputStream.ensureFill(RedisInputStream.java:205)
        at redis.clients.jedis.util.RedisInputStream.readByte(RedisInputStream.java:43)
        at redis.clients.jedis.Protocol.process(Protocol.java:158)
        at redis.clients.jedis.Protocol.read(Protocol.java:223)
        at redis.clients.jedis.Connection.readProtocolWithCheckingBroken(Connection.java:352)
        at redis.clients.jedis.Connection.getStatusCodeReply(Connection.java:270)
        at redis.clients.jedis.BinaryJedis.ping(BinaryJedis.java:381)
        at com.github.anigkus.redistls.RedisDemo.main(RedisDemo.java:9)
Caused by: javax.net.ssl.SSLHandshakeException: Received fatal alert: certificate_required
        at java.base/sun.security.ssl.Alert.createSSLException(Alert.java:131)
        at java.base/sun.security.ssl.Alert.createSSLException(Alert.java:117)
        at java.base/sun.security.ssl.TransportContext.fatal(TransportContext.java:314)
        at java.base/sun.security.ssl.Alert$AlertConsumer.consume(Alert.java:293)
        at java.base/sun.security.ssl.TransportContext.dispatch(TransportContext.java:187)
        at java.base/sun.security.ssl.SSLTransport.decode(SSLTransport.java:164)
        at java.base/sun.security.ssl.SSLSocketImpl.decode(SSLSocketImpl.java:1154)
        at java.base/sun.security.ssl.SSLSocketImpl.readApplicationRecord(SSLSocketImpl.java:1124)
        at java.base/sun.security.ssl.SSLSocketImpl$AppInputStream.read(SSLSocketImpl.java:823)
        at java.base/java.io.InputStream.read(InputStream.java:205)
        at redis.clients.jedis.util.RedisInputStream.ensureFill(RedisInputStream.java:199)
        ... 7 more
```

## Resolve E1:

<mark> Remove client certificate verification logic </mark>

```
127.0.0.1:6379> CONFIG  get tls-auth-clients
1) "tls-auth-clients"
2) "yes"
127.0.0.1:6379> CONFIG  set tls-auth-clients no
OK
127.0.0.1:6379> CONFIG  get tls-auth-clients
1) "tls-auth-clients"
2) "no"
```

## E2

```
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
Exception in thread "main" redis.clients.jedis.exceptions.JedisConnectionException: Could not get a resource from the pool
        at redis.clients.jedis.util.Pool.getResource(Pool.java:84)
        at redis.clients.jedis.JedisPool.getResource(JedisPool.java:366)
        at com.github.anigkus.redistls.JedisSSLTest.main(JedisSSLTest.java:45)
Caused by: redis.clients.jedis.exceptions.JedisDataException: DENIED Redis is running in protected mode because protected mode is enabled, no bind address was specified, no authentication password is requested to clients. In this mode connections are only accepted from the loopback interface. If you want to connect from external computers to Redis you may adopt one of the following solutions: 1) Just disable protected mode sending the command 'CONFIG SET protected-mode no' from the loopback interface by connecting to Redis from the same host the server is running, however MAKE SURE Redis is not publicly accessible from internet if you do so. Use CONFIG REWRITE to make this change permanent. 2) Alternatively you can just disable the protected mode by editing the Redis configuration file, and setting the protected mode option to 'no', and then restarting the server. 3) If you started the server manually just for testing, restart it with the '--protected-mode no' option. 4) Setup a bind address or an authentication password. NOTE: You only need to do one of the above things in order for the server to start accepting connections from the outside.
        at redis.clients.jedis.Protocol.processError(Protocol.java:135)
        at redis.clients.jedis.Protocol.process(Protocol.java:169)
        at redis.clients.jedis.Protocol.read(Protocol.java:223)
        at redis.clients.jedis.Connection.readProtocolWithCheckingBroken(Connection.java:352)
        at redis.clients.jedis.Connection.getStatusCodeReply(Connection.java:270)
        at redis.clients.jedis.BinaryJedis.auth(BinaryJedis.java:2717)
        at redis.clients.jedis.BinaryJedis.initializeFromClientConfig(BinaryJedis.java:94)
        at redis.clients.jedis.BinaryJedis.<init>(BinaryJedis.java:292)
        at redis.clients.jedis.Jedis.<init>(Jedis.java:167)
        at redis.clients.jedis.JedisFactory.makeObject(JedisFactory.java:177)
        at org.apache.commons.pool2.impl.GenericObjectPool.create(GenericObjectPool.java:918)
        at org.apache.commons.pool2.impl.GenericObjectPool.borrowObject(GenericObjectPool.java:431)
        at org.apache.commons.pool2.impl.GenericObjectPool.borrowObject(GenericObjectPool.java:356)
        at redis.clients.jedis.util.Pool.getResource(Pool.java:75)
        ... 2 more
```

Resolve E2:

<mark> remove protected mode </mark> 

```
127.0.0.1:6379> CONFIG GET protected-mode
1) "protected-mode"
2) "yes"
127.0.0.1:6379> CONFIG set protected-mode no 
OK
127.0.0.1:6379> CONFIG GET protected-mode
1) "protected-mode"
2) "no"
```


## E3

```
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
Exception in thread "main" redis.clients.jedis.exceptions.JedisConnectionException: Could not get a resource from the pool
        at redis.clients.jedis.util.Pool.getResource(Pool.java:84)
        at redis.clients.jedis.JedisPool.getResource(JedisPool.java:366)
        at com.github.anigkus.redistls.JedisSSLTest.main(JedisSSLTest.java:45)
Caused by: redis.clients.jedis.exceptions.JedisDataException: ERR AUTH <password> called without any password configured for the default user. Are you sure your configuration is correct?
        at redis.clients.jedis.Protocol.processError(Protocol.java:135)
        at redis.clients.jedis.Protocol.process(Protocol.java:169)
        at redis.clients.jedis.Protocol.read(Protocol.java:223)
        at redis.clients.jedis.Connection.readProtocolWithCheckingBroken(Connection.java:352)
        at redis.clients.jedis.Connection.getStatusCodeReply(Connection.java:270)
        at redis.clients.jedis.BinaryJedis.auth(BinaryJedis.java:2717)
        at redis.clients.jedis.BinaryJedis.initializeFromClientConfig(BinaryJedis.java:94)
        at redis.clients.jedis.BinaryJedis.<init>(BinaryJedis.java:292)
        at redis.clients.jedis.Jedis.<init>(Jedis.java:167)
        at redis.clients.jedis.JedisFactory.makeObject(JedisFactory.java:177)
        at org.apache.commons.pool2.impl.GenericObjectPool.create(GenericObjectPool.java:918)
        at org.apache.commons.pool2.impl.GenericObjectPool.borrowObject(GenericObjectPool.java:431)
        at org.apache.commons.pool2.impl.GenericObjectPool.borrowObject(GenericObjectPool.java:356)
        at redis.clients.jedis.util.Pool.getResource(Pool.java:75)
        ... 2 more
```

Resolve E3:

<mark> The server does not have a password configured, so you need to remove the client password. </mark> 


# Success
```
127.0.0.1:6379> KEYS *
1) "key"
2) "key1"
```