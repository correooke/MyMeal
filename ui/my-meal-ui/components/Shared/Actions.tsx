import React from "react";
import IsActiveOption from "./IsActiveOption";
import DeleteButton from "./DeleteButton";

const Actions: React.FC = () => {
  return (
    <>
      <DeleteButton isPrimaryColor={true} onClick={() => {}} />
      <IsActiveOption
        isPrimaryColor
        checked={true}
        onChange={(checked: boolean) => {}}
      />
    </>
  );
};

export default Actions;
