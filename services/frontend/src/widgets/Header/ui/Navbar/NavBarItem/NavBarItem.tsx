import { Link } from "react-router-dom";
import styles from "./NavBarItem.module.scss";
import clsx from "clsx";

interface NavigationItemProps {
  item: string;
  index: number;
  isActive: boolean;
  onItemClicked: (index: number) => void;
  backgroundImage: string;
  activeIndex: number;
}

export const NavBarItem: React.FC<NavigationItemProps> = ({
  item,
  index,
  isActive,
  activeIndex,
  onItemClicked,
}) => {
  const activeDecorationClass = (index: number) => {
    switch (index) {
      case 0:
        return styles.decoration_active_chair;
      case 1:
        return styles.decoration_active_services;
      case 2:
        return styles.decoration_active_play;
      default:
        return "";
    }
  };
  return (
    <li
      className={clsx(styles.list__item)}
      onClick={() => onItemClicked(index)}
    >
      <div className={styles.decoration_container}>
        <div
          className={clsx(styles.decoration, {
            [styles.decoration_active]: index === activeIndex,
            [activeDecorationClass(index)]: index === activeIndex,
          })}
        ></div>
      </div>
      <Link
        to="/"
        className={clsx(styles.link, { [styles.link_active]: isActive })}
      >
        {item}
      </Link>
    </li>
  );
};
