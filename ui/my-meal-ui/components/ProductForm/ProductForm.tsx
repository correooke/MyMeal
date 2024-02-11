"use client";

import { Product } from "@/models/product";
import { Box } from "@mui/material";
import { FC, useState } from "react";
import ProductInput from "./components/ProductInput";
import ProductDescriptionInput from "./components/ProductDescriptionInput";
import PriceInput from "./components/PriceInput";
import DeleteButton from "./components/DeleteButton";
import IsActiveOption from "./components/IsActionOption";
import ProductDetails from "./components/ProductDetails";
import ProductImageUploader from "./components/ProductImageUploader";

interface ProductFormProps {
  product: Product;
}

const ProductForm: FC<ProductFormProps> = ({ product }) => {
  const [updatedProduct, setProduct] = useState({
    ...product,
  });

  const handleChange = (value: string | number, name: string) => {
    setProduct({
      ...updatedProduct,
      [name]: value,
    });
  };

  return (
    <Box sx={{ maxWidth: "650px" }}>
      <ProductImageUploader />
      <ProductInput value={updatedProduct.name} onChange={handleChange} />
      <IsActiveOption
        checked={updatedProduct.isActive}
        onChange={(checked: boolean) => {}}
      />
      <DeleteButton onClick={() => {}} />
      <ProductDescriptionInput
        value={updatedProduct.description}
        onChange={handleChange}
      />
      <ProductDetails product={updatedProduct} />
      <PriceInput value={updatedProduct.price} onChange={handleChange} />
    </Box>
  );
};

export default ProductForm;
