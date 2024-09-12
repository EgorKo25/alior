import axios from "axios";

export interface ProjData {
  title: string;
  description: string;
  url: string;
}

export const getProjects = async () => {
  return axios
    .get("http://localhost:8000/projects")
    .then((response) => {
      return response.data as ProjData[];
    })
    .catch((error) => {
      console.log(error);
      throw new Error(error);
    });
};
