import { RouterProvider } from "react-router-dom";
import { router } from "./Router";
import "./outputTW.css";

export const App = () => {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
};
