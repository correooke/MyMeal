import { TextField } from "@mui/material";
import { NumberProps } from "../../Shared/NumberProps";

const PriceInput = ({ value, onChange }: NumberProps) => (
  <TextField
    name={"price"}
    value={value}
    label={"Precio"}
    onChange={(event) =>
      onChange(parseInt(event.target.value), event.target.name)
    }
    fullWidth
    type="number"
  />
);

export default PriceInput;
