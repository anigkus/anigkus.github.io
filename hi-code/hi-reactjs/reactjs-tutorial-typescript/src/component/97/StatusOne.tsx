type statusOneProps = {
    status:'success'|'error'|'loading'
}


const StatusOne = (props: statusOneProps) => {
    let message = '';
    if( props.status ==='success'){
        message="Loaded Success"
    }
    if( props.status ==='error'){
        message="Loaded Error"
    }
    if( props.status ==='loading'){
        message="Loading"
    }

  return (
    <div>
      {message}
    </div>
  );
}

export default StatusOne;
