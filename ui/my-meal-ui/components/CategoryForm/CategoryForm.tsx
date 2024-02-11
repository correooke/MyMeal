import React from "react";
import Box from "@mui/material/Box";
import ProductForm from "../ProductForm/ProductForm";
import { Product } from "@/models/product";

const CategoryForm = () => {
  const mockProduct: Product = {
    name: "Sample Product",
    price: 10.99,
    description: "This is a sample product description",
    id: 0,
    imageUrl: "",
    isActive: false,
  };
  return (
    <Box display={"flex"} sx={{ justifyContent: "center", width: "100%" }}>
      <ProductForm product={mockProduct} />
    </Box>
  );
};

export default CategoryForm;
