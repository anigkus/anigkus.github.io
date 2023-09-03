type ListProps<T> = {
  items: T[];
  onClick: (value: T) => void;
};

export const ListOne = <T extends {}>({ items, onClick }: ListProps<T>) => {
  function objectToQueryString(obj: any) {
    const keys = Object.keys(obj);
    const keyValuePairs = keys.map((key) => {
      //return encodeURIComponent(key) + '=' + encodeURIComponent(obj[key]);
      return obj[key];
    });
    return keyValuePairs;
    //return keyValuePairs.join('&');
  }
  return (
    <div>
      <h2>ListOne item map</h2>
      <br />
      {items.map((item, index) => {
        // return (
        //   <div>
        //     {' '}
        //     <div></div>
        //   </div>
        // );

        return (
          <button key={index} onClick={() => onClick(item)}>
            {/* {JSON.strin gify(item)} */}
            {/* {objectToQueryString(item)}  */}xxx
          </button>
        );
      })}
    </div>
  );
};

export const ListTwo = <T extends { id: number }>({
  items,
  onClick
}: ListProps<T>) => {
  return (
    <div>
      <h2>ListTwo item map</h2>
      <br />
      {items.map((item, index) => {
        return (
          <button key={index} onClick={() => onClick(item)}>
            {item.id}
          </button>
        );
      })}
    </div>
  );
};

export const ListThree = <T extends string | number>({
  items,
  onClick
}: ListProps<T>) => {
  return (
    <div>
      <h2>ListThree item map</h2>
      <br />
      {items.map((item, index) => {
        return (
          <button key={index} onClick={() => onClick(item)}>
            {item}
          </button>
        );
      })}
    </div>
  );
};
