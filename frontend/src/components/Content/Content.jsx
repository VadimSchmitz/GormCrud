import MovieEntry from "./MovieEntry/MovieEntry";

export default function Content({ movies }) {
  return (
    <div>
      {movies.map((movie, index) => (
        <MovieEntry
          key={index}
          id={index}
          title={movie.title}
          rating={movie.rating}
          year={movie.year}
          summary={movie.summary}
        />
      ))}
    </div>
  );
}
