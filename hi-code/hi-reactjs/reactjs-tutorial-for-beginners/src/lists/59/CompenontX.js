import React from 'react';
import CompenontY from './CompenontY';

export const NameContext = React.createContext();
export const ChannerContext = React.createContext();

export default function CompenontX() {
  return (
    <div>
      <NameContext.Provider value="Anigkus">
        <ChannerContext.Provider value="Channel">
          <CompenontY />
        </ChannerContext.Provider>
      </NameContext.Provider>
    </div>
  );
}
