import { Formik, Form } from "formik";
import FormikField from "./FormikField";

export default function MovieForm({ setMovies }) {
  return (
    <div>
      <h1>Add a movie</h1>
      <Formik
        initialValues={{
          imdb_id: "",
          title: "",
          rating: "",
          year: "",
          summary: ""
        }}
        onSubmit={async (values) => {
          await new Promise((r) => setTimeout(r, 20));
          const requestOptions = {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(values)
          };
          console.log(requestOptions.body);
          fetch("http://localhost:8090/movies", requestOptions).then((response) => response.json());

          setMovies((movies) => [...movies, values]);
        }}>
        <Form className="flex-col">
          <FormikField id="imdb_id" label="IMDb_id" placeholder="tt0993846" type="" />
          <FormikField id="title" label="Title" placeholder="The Wolf of Wall Street" type="" />
          <FormikField id="rating" label="Rating" placeholder="8.2" type="number" />
          <FormikField id="year" label="Year" placeholder="2013" type="number" />
          <FormikField
            id="summary"
            label="Summary"
            placeholder="Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government."
            type=""
          />
          <button
            className="max-w-[160px] py-2 px-3 bg-blue-500 hover:bg-blue-700 text-white font-bold"
            type="submit">
            Add a movie
          </button>
        </Form>
      </Formik>
    </div>
  );
}
