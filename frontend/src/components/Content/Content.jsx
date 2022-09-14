import MovieEntry from "./MovieEntry/MovieEntry";
import "./Content.css";

export default function Content({ movies }) {
  console.log(movies);
  return (
    <div>
      <MovieEntry
        id="5"
        title="The Chronicles of Narnia: The Lion, the Witch and the Wardrobe"
        year="2005"
        rating="10"
        plot="Four kids travel through a wardrobe to the land of Narnia and learn of their destiny to free it with the guidance of a mystical lion."
      />
      <MovieEntry
        id="1"
        title="The Shawshank Redemption"
        year="1993"
        rating="9.3"
        plot="Two imprisoned men bond over a number of years, finding solace and eventual redemption
              through acts of common decency."
      />
      <MovieEntry
        id="2"
        title="American Gangster"
        year="2007"
        rating="6.3"
        plot="An outcast New York City cop is charged with bringing down Harlem drug lord Frank Lucas, whose real life inspired this partly biographical film."
      />
      <MovieEntry
        id="3"
        title="Pulp Fiction"
        year="1994"
        rating="10"
        plot="The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption."
      />
    </div>
  );
}
