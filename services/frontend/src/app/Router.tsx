import { createBrowserRouter } from "react-router-dom";
import { HomePage } from "src/pages/HomePage";
import { Layout } from "./Layout";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        index: true,
        element: <HomePage />,
      },
    ],
  },
]);
