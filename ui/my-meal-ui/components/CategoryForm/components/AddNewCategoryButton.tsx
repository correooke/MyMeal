import React from "react";
import { Button, Typography } from "@mui/material";

const AddNewCategoryButton: React.FC = () => {
  return (
    <Button variant="contained" sx={{ bgcolor: "primary.dark" }}>
      Agregar nueva categoría
    </Button>
  );
};

export default AddNewCategoryButton;
