import React, { ChangeEvent } from "react";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import { Button } from "@mui/material";

type IsActiveOptionProps = {
  checked: boolean;
  onChange: (checked: boolean) => void;
  isPrimaryColor: boolean;
};
const IsVisibleIcon = (checked: boolean, isPrimaryColor: boolean) =>
  checked ? (
    <VisibilityIcon color={isPrimaryColor ? "primary" : "secondary"} />
  ) : (
    <VisibilityOffIcon color={isPrimaryColor ? "primary" : "secondary"} />
  );

const IsActiveOption = ({
  checked,
  isPrimaryColor,
  onChange,
}: IsActiveOptionProps) => (
  <Button
    startIcon={IsVisibleIcon(checked, isPrimaryColor)}
    onClick={() => onChange(checked)}
    variant="text"
    color={isPrimaryColor ? "primary" : "secondary"}
  >
    Ocultar
  </Button>
);

export default IsActiveOption;
