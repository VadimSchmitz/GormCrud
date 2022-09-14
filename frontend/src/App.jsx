import { useEffect, useState } from "react";
import Content from "./components/Content/Content";
import MovieForm from "./components/AddMovieForm/MovieForm";
import "./App.css";

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
    <div className="content">
      <h1>Movies in the IMDb database</h1>
      <p>This is a list of all the movies currently stored in the database</p>
      <hr></hr>
      <MovieForm />
      <hr></hr>
      <Content movies={movies} />
    </div>
  );
}

export default App;
