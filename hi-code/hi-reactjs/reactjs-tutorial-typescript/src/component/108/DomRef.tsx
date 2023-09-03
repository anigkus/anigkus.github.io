import { useRef,useEffect } from "react"

export const DomRef = ()  =>{
    const inputRef = useRef<HTMLInputElement>(null!)

    useEffect(() => {
        inputRef.current.focus()
      return () => {
       
      };
    }, []);

return <input ref={inputRef} defaultValue="Default..." type="input"></input>
}