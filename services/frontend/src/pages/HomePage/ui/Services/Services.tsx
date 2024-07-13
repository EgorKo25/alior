import { MainButton } from "src/shared/ui/MainButton";
import { ServiceCard } from "./ui/ServiceCard";
import styles from "./Servicies.module.scss";
import { useEffect, useState } from "react";
export const Services = () => {
  const [activeCard, setActiveCard] = useState<number | null>(null);

  const handleScroll = () => {
    const cards = document.querySelectorAll("[data-card-index]");
    let newActiveCard = null;

    cards.forEach((card, index) => {
      const cardRect = card.getBoundingClientRect();
      const isFullyVisible =
        cardRect.top >= 0 && cardRect.bottom <= window.innerHeight;

      if (isFullyVisible) {
        newActiveCard = index;
      }
    });

    setActiveCard(newActiveCard);
  };

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  return (
    <section id="services" className={styles.service_container}>
      <div className={styles.service_cards_container}>
        {[0, 1, 2].map((index) => (
          <ServiceCard
            key={index}
            index={index}
            isActive={index === activeCard}
          />
        ))}
      </div>
      <div className={styles.service_heading_container}>
        <div className={styles.heading_container}>
          <h2 className={styles.service_heading}>
            Что мы <br />{" "}
            <span className={styles.service_heading_span}>предлагаем</span>
          </h2>
          <div className={styles.parallelogram}></div>
        </div>
        <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-accent-white"}
        />
        {/* <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-white-accent"}
        />
        <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-orange-white"}
        />
        <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-blue-white"}
        />
        <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-black-white"}
        />
        <MainButton
          className={styles.button}
          title="Уже знаю, что хочу"
          colorSchema={" btn-white-black"}
        /> */}
      </div>
    </section>
  );
};
