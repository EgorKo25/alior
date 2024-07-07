export const Message: React.FC<{ title: string }> = ({ title }) => {
  return (
    <div
      className={` block border-accent border-solid border p-2 pt-2 pb-2 pl-4 pr-5 font-light text-sm lg:text-base rounded-20 rounded-bl-none max-w-96 h-fit`}
    >
      {title}
    </div>
  );
};
