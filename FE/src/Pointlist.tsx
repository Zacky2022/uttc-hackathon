import { Link } from "react-router-dom";

const List = () => {
  return (
    <>
      <h1>ポイント一覧</h1>
      <div>
        <Link to={`/`}>Login画面に戻る</Link>
      </div>
    </>
  );
};

export default List;