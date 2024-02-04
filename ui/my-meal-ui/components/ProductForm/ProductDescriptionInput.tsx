import React from "react";
import { TextField } from "@mui/material";
import { InputProps } from "./StringProps";

const ProductDescriptionInput = ({ value, onChange }: InputProps) => (
  <TextField
    name={"product-description"}
    value={value}
    onChange={(event) => onChange(event.target.value, event.target.name)}
    label="Description"
    multiline
    rows={4}
    fullWidth
    margin="normal"
  />
);

export default ProductDescriptionInput;