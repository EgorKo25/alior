import { useEffect, useState } from "react";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { ProjData, projectsData } from "src/entities/projects/";
import { Slide } from "./ui/Slide";
import "./styles.css";

const SLIDE_BG_GRADIENTS_CONFIG: {
  colorSchema: "orange" | "blue" | "black";
  bg: string;
}[] = [
  { colorSchema: "orange", bg: "via-orange-100/60" },
  { colorSchema: "blue", bg: "via-blue-100/60" },
  { colorSchema: "black", bg: "via-black/30" },
];

export const Carousel = () => {
  const [projects, setProjects] = useState<ProjData[]>([]);
  const [isProjectsLoaded, setIsProjectsLoaded] = useState(false);
  const [activeSlide, setActiveSlide] = useState(0);

  useEffect(() => {
    const loadProjects = async () => {
      setProjects(await projectsData());
      setIsProjectsLoaded(true);
    };
    loadProjects();
  }, []);

  let settings = {
    dots: false,
    infinite: false,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    nextArrow: <></>,
    prevArrow: <></>,
    beforeChange: (_: number, next: number) => {
      setActiveSlide(next);
    },
  };
  return isProjectsLoaded && projects.length > 0 ? (
    <section
      id="cases"
      className={`w-screen overflow-hidden transition-all bg-gradient-to-b from-transparent ${
        SLIDE_BG_GRADIENTS_CONFIG[activeSlide % 3].bg
      } via-[percentage:20%_70%] to-transparent`}
    >
      <div className="slider-wrapper mt-28 mb-40 md:mx-[30px] xl:mx-24">
        <Slider {...settings}>
          {projects.map((project, index) => {
            return (
              <Slide
                ColorSchema={SLIDE_BG_GRADIENTS_CONFIG[index % 3].colorSchema}
                title={project.title}
                description={project.description}
                projectUrl={project.url}
                key={index}
                isActive={index === activeSlide}
              />
            );
          })}
        </Slider>
      </div>
    </section>
  ) : (
    <></>
  );
};
