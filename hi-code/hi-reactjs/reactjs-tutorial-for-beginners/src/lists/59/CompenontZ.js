import React from 'react';
import { ChannerContext, NameContext } from './CompenontX';

export default function CompenontZ() {
  return (
    <div>
      <NameContext.Consumer>
        {(nameValue) => {
          return (
            <>
              <ChannerContext.Consumer>
                {(channerValue) => {
                  return (
                    <div>
                      NameContext :{nameValue},ChannerContext : {channerValue}
                    </div>
                  );
                }}
              </ChannerContext.Consumer>
            </>
          );
        }}
      </NameContext.Consumer>
    </div>
  );
}
