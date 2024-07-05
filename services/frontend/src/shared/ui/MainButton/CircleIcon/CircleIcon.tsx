import clsx from "clsx";
import styles from "../MainButton.module.scss";

export const CircleIcon: React.FC<{ className?: string }> = ({ className }) => (
  <div className={clsx(styles.circle_container)}>
    <div className={clsx(styles.circle, className)}></div>
  </div>
);
