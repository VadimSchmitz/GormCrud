import { Field } from "formik";

export default function FormikField({ id, label, placeholder, type, error, touched }) {
  return (
    <div className="mb-3 ">
      <label className="block text-gray-700 text-sm font-bold mb-[1px]" htmlFor={id}>
        {label}
      </label>
      <div className="flex">
        <Field
          className="shadow appearance-none border py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id={id}
          name={id}
          placeholder={placeholder}
          type={type}
        />
      </div>
      {error && touched ? <p className="my-auto text-red-600">{error}</p> : null}
    </div>
  );
}
