import { useEffect } from 'react';

export default function useDocumenTitle(count) {
  useEffect(() => {
    document.title = `count ${count}`;
  }, [count]);
}
