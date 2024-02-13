import { TextField } from "@mui/material";
import { InputProps } from "../../Shared/StringProps";

const CategoryInput = ({ value, onChange }: InputProps) => (
  <TextField
    name={"category"}
    value={value}
    label={"Categoría"}
    onChange={(event) => onChange(event.target.value, event.target.name)}
    sx={{ width: "450px", marginLeft: "auto", marginRight: "20px" }}
    color="primary"
  />
);

export default CategoryInput;
