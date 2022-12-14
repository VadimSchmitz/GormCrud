import MovieEntry from "./MovieEntry";

const PaginatedMovies = ({ data }) => {
  return (
    <div>
      {data.map((movie, index) => (
        <MovieEntry
          key={index}
          id={movie.imdb_id}
          title={movie.title}
          rating={movie.rating}
          year={movie.year}
          summary={movie.summary}
        />
      ))}
    </div>
  );
};

export default PaginatedMovies;
