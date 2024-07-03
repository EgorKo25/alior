import { TgIcon } from "src/shared/ui/TgButton/TgIcon";
import { QuestionIcon } from "./ui/QuestionIcon";
import styles from "./Footer.module.scss";
import clsx from "clsx";
export const Footer = () => {
  return (
    <footer className={clsx(styles.footer)}>
      <div className={clsx(styles.footer_heading_container)}>
        <h2 className={clsx(styles.footer_heading)}>Alior</h2>
        <p className={clsx(styles.footer_decoration_paragraph)}>
          От вас идея — от нас реализация!
        </p>
      </div>
      <div className={clsx(styles.bottom_footer_container)}>
        <div className={clsx(styles.policy_container)}>
          <h4>©2024 Alior</h4>
          <h4>Политика обработки данных</h4>
        </div>
        <div className={clsx(styles.policy_button_container)}>
          <button className={clsx(styles.policy_button)}>
            <QuestionIcon color="#fff" size={24} />
          </button>
          <button className={clsx(styles.policy_button)}>
            <TgIcon color="#fff" size={24} />
          </button>
        </div>
      </div>
    </footer>
  );
};
