import { Formik, Form } from "formik";
import FormikField from "./FormikField";

export default function MovieForm() {
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
          await new Promise((r) => setTimeout(r, 500));
          console.log(JSON.stringify(values, null, 2));
        }}>
        <Form className="flex-col">
          <FormikField id="imdb_id" label="IMDb_id" placeholder="231" type="" />
          <FormikField id="title" label="Title" placeholder="narnia" type="" />
          <div className="flex space-x-3">
            <FormikField id="rating" label="Rating" placeholder="9.1" type="" />
            <FormikField id="year" label="Year" placeholder="1999" type="" />
          </div>

          <FormikField id="summary" label="Summary" placeholder="wee wee woo woo" type="" />
          <button
            className="max-w-[160px] py-2 px-3 bg-blue-500 hover:bg-blue-700 text-white font-bold"
            type="submit">
            Add movie
          </button>
        </Form>
      </Formik>
    </div>
  );
}
