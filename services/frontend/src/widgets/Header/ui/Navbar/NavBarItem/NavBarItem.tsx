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
  activeIndex,
  isActive,
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
            " bg-black sm:bg-black md:bg-black lg:bg-white": !isActive,
            " bg-transparent sm:fill-black md:fill-black": isActive,
            [styles.decoration_active]: index === activeIndex,
            [activeDecorationClass(index)]: index === activeIndex,
          })}
        ></div>
      </div>
      <Link
        to="/"
        className={
          clsx(styles.link) + " sm:text-black md:text-black lg:text-white"
        }
      >
        {item}
      </Link>
    </li>
  );
};
