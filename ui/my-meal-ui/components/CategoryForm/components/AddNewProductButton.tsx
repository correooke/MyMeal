import { FC } from "react";
import Button from "@mui/material/Button";

type AddNewProductButtonProps = {
  onClick: () => void;
};
const AddNewProductButton: FC<AddNewProductButtonProps> = ({ onClick }) => {
  const handleAddProduct = () => {
    onClick();
  };

  return (
    <Button variant="contained" onClick={handleAddProduct}>
      Agregar otro producto
    </Button>
  );
};

export default AddNewProductButton;
