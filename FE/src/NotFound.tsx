import { Link } from "react-router-dom";

const NotFound = () => {
  return (
    <>
      <h1>Page Not Found</h1>
      <div>
        <Link to={`/`}>Go back to Lorgin Form</Link>
      </div>
    </>
  );
};

export default NotFound;