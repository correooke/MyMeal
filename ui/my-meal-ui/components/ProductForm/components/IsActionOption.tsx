import React, { ChangeEvent } from "react";
import { Switch, FormControlLabel } from "@mui/material";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";

type IsActiveOptionProps = {
  checked: boolean;
  onChange: (checked: boolean) => void;
};

const IsActiveOption = ({ checked, onChange }: IsActiveOptionProps) => (
  <FormControlLabel
    control={
      <Switch
        checked={checked}
        onChange={(event: ChangeEvent<HTMLInputElement>) =>
          onChange(event.target.checked)
        }
        name="isActive"
      />
    }
    label={checked ? <VisibilityIcon /> : <VisibilityOffIcon />}
    labelPlacement="start"
  />
);

export default IsActiveOption;
