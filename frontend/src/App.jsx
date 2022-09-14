import { useEffect, useState } from "react";
import Content from "./components/Content/Content";
import MovieForm from "./components/AddMovieForm/MovieForm";

function App() {
  const [movies, setMovies] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8090/movies")
      .then((res) => res.json())
      .then((data) => {
        setMovies(data);
      });
  }, []);

  return (
    <div className="content mx-auto bg-[#F8F8F8] p-5 w-9/12 mt-10 mb-16">
      <h1>Movies in the IMDb database</h1>
      <p>This is a list of all the movies currently stored in the database</p>
      <hr></hr>
      <MovieForm setMovies={setMovies} />
      <hr></hr>
      <Content movies={movies} />
    </div>
  );
}

export default App;
