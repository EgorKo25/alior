import { MainButton } from "src/shared/ui/MainButton";
import { ServiceCard } from "./ui/ServiceCard";
import styles from "./Servicies.module.scss";
import { useEffect, useState } from "react";
import { HashLink } from "react-router-hash-link";
export const Services = () => {
  const [activeCard, setActiveCard] = useState<number | null>(null);

  const handleScroll = () => {
    const cards = document.querySelectorAll("[data-card-index]");

    cards.forEach((card, index) => {
      const cardRect = card.getBoundingClientRect();
      const isMiddle =
        cardRect.top <= window.innerHeight / 2 &&
        cardRect.bottom >= window.innerHeight / 2 - 17;

      if (isMiddle && index != activeCard) {
        setActiveCard(index);
      }
    });
  };

  useEffect(() => {
    setActiveCard(0);
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
        <HashLink smooth to={"#consult"}>
          <MainButton
            className={styles.button}
            title="Уже знаю, что хочу"
            colorSchema={" btn-accent-white"}
          />
        </HashLink>
      </div>
    </section>
  );
};
