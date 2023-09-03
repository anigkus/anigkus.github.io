import React, { Component } from 'react';

type ClassProps = {
    message:string
}

type StateProps = {
    count:number
}
export class ClassCounter extends Component<ClassProps,StateProps> {
   state = {
         count:0
      }
    

    handlerClick = () => {
        this.setState((prevState) => ({count:prevState.count+1}))
    } 
  render() {
    return (
      <div>
        <button onClick={this.handlerClick}>Click</button>
        {this.props.message},{this.state.count}
      </div>
    );
  }
}

export default ClassCounter;
