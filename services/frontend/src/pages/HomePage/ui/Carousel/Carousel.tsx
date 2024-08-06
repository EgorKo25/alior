import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { Slide } from "./Slide";
import "./styles.css";
import { useState } from "react";

const SLIDE_BG_GRADIENTS_CONFIG = [
  "via-orange-100/60",
  "via-blue-100/60",
  "via-black/30",
];
export const Carousel = () => {
  const [activeSlide, setActiveSlide] = useState(0);

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
  return (
    <section
      className={`w-screen overflow-hidden transition-all bg-gradient-to-b from-transparent ${SLIDE_BG_GRADIENTS_CONFIG[activeSlide]} via-[percentage:20%_70%] to-transparent`}
    >
      <div className="slider-wrapper mt-28 mb-40 md:mx-[30px] xl:mx-24">
        <Slider {...settings}>
          <Slide
            ColorSchema={"orange"}
            title={"Unicorn bot"}
            description={
              "Бот для управления бронированием в ресторане, который интегрируется с CRM-системой клиента, автоматизируя процессы и снижая нагрузку на персонал"
            }
            projectUrl=""
          />
          <Slide
            ColorSchema={"blue"}
            title={"Get магазин"}
            description={
              "Сайт с конфигуратором, для мебельной компании. Он позволяет кастомизировать мебель, выбирая материалы, цвета и размеры. Это увеличило их продажи и количество довольных клиентов"
            }
            projectUrl=""
          />
          <Slide
            ColorSchema={"black"}
            title={"Landing page"}
            description={
              "Сайт с конфигуратором, для мебельной компании. Он позволяет кастомизировать мебель, выбирая материалы, цвета и размеры. Это увеличило их продажи и количество довольных клиентов"
            }
            projectUrl=""
          />
        </Slider>
      </div>
    </section>
  );
};
