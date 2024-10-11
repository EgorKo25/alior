import clsx from "clsx";
import { QuestionIcon } from "./ui/QuestionIcon";
import { GitIcon } from "./ui/GitIcon";
import { TgFooterIcon } from "./ui/TgFooterIcon";
import { LinkedinIcon } from "./ui/LinkedinIcon";
import styles from "./Footer.module.scss";

export const Footer = () => {
  return (
    <footer
      className={
        clsx(styles.footer) +
        ` pt-4 pb-8 lg:pb-12 xl:pb-16  lg:h-screen lg:max-h-[65vw] flex flex-col justify-around`
      }
    >
      <div
        className={clsx(
          styles.footer_heading_container + ` mb-8 sm:mb-12 lg:mb-8 xl:mb-0`
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
      <div className={clsx(styles.bottom_footer_container) + ` lg:text-2xl`}>
        <div className={clsx(styles.policy_container)}>
          <h4>©2024 Alior</h4>
          <h4>Политика обработки данных</h4>
        </div>
        <div className={clsx(styles.policy_button_container)}>
          <button
            className={
              " flex justify-center items-center size-4 md:size-6 lg:size-8"
            }
          >
            <QuestionIcon className=" hover:text-accent" />
          </button>
          <button
            className={
              " flex justify-center items-center size-4 md:size-6 lg:size-8"
            }
          >
            <LinkedinIcon className={" hover:text-accent"} />
          </button>
          <button
            className={
              " flex justify-center items-center size-4 md:size-6 lg:size-8"
            }
          >
            <GitIcon className={" hover:text-accent"} />
          </button>
          <button
            className={
              " flex justify-center items-center size-4 md:size-6 lg:size-8"
            }
          >
            <TgFooterIcon className={" hover:text-accent"} />
          </button>
        </div>
      </div>
    </footer>
  );
};
