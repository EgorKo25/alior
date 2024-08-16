import clsx from "clsx";
import { QuestionIcon } from "./ui/QuestionIcon";
import { GitIcon } from "./ui/GitIcon";
import { TgFooterIcon } from "./ui/TgFooterIcon";
import { LinkedinIcon } from "./ui/LinkedinIcon";
import styles from "./Footer.module.scss";

export const Footer = () => {
  return (
    <footer className={clsx(styles.footer)}>
      <div
        className={clsx(
          styles.footer_heading_container + ` mb-8 sm:mb-12 lg:mb-16`
        )}
      >
        <h2
          className={clsx(styles.footer_heading) + ` font-bold leading-[1.1]`}
        >
          Alior
        </h2>
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
          <button className={" flex justify-center items-center size-6"}>
            <QuestionIcon className=" hover:text-accent" />
          </button>
          <button className={" flex justify-center items-center size-6"}>
            <LinkedinIcon className={" hover:text-accent"} />
          </button>
          <button className={" flex justify-center items-center size-6"}>
            <GitIcon className={" hover:text-accent"} />
          </button>
          <button className={" flex justify-center items-center size-6"}>
            <TgFooterIcon className={" hover:text-accent"} />
          </button>
        </div>
      </div>
    </footer>
  );
};
