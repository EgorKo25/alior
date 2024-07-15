export const CustomRadioInput: React.FC<{
  name: string;
  label: string;
  id: string;
  isChecked: boolean;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
}> = ({ name, label, id, isChecked, onChange }) => {
  return (
    <div className="one-option">
      <input
        type="radio"
        name={name}
        id={id}
        className=" absolute -z-10 opacity-0"
        checked={isChecked}
        onChange={onChange}
      />
      <label
        htmlFor={id}
        className="text-white ml-11 mr-4 text-sm sm:text-base relative"
      >
        <div
          className={`new-radio absolute -left-7 top-0 w-4 h-4 border-2 border-white bg-accent rounded-full`}
        >
          <div
            className={`inside-circle w-2 h-2 ml-0.5 mt-0.5 bg-accent rounded-full transition-all ${
              isChecked ? "bg-white" : ""
            }`}
          ></div>
        </div>
        {label}
      </label>
    </div>
  );
};
