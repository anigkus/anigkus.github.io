export type ProfileOneProps= {
    name:string
}
export const ProfileOne = ({name}: ProfileOneProps)=> {
    return <div>Private Profile of {name}</div>
}