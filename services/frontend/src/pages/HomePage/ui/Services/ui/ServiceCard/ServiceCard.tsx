import { ServiceCardsData } from "src/features/appData";
import { MainButton } from "src/shared/ui/MainButton";
import styles from "./ServicieCard.module.scss";
import clsx from "clsx";
import { HashLink } from "react-router-hash-link";
export const ServiceCard = ({
  index,
  isActive,
}: {
  index: number;
  isActive: boolean;
}) => {
  return (
    <div
      className={clsx(styles.card_container, {
        [styles.card_active]: isActive,
      })}
      data-card-index={index}
    >
      <h3 className={styles.card_title}>{ServiceCardsData[index].title}</h3>
      <p className={styles.card_discription}>
        {ServiceCardsData[index].description}
      </p>
      {isActive && (
        <HashLink smooth to={"#consult"}>
          <MainButton
            colorSchema=" btn-orange-white"
            className={styles.button}
            title="Мне это надо"
          />
        </HashLink>
      )}
    </div>
  );
};
