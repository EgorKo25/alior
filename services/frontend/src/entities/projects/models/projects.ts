import { ProjectsDummyData } from "src/features/appData";
import { getProjects } from "../api/getProjects";

export const projectsData = async () => {
  try {
    const projects = await getProjects();
    return projects;
  } catch (error) {
    console.error(
      "Проекты не доступны по api, вместо этого подставлены заглушки."
    );
    return ProjectsDummyData;
  }
};
