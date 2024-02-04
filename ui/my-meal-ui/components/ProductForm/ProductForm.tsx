"use client";

import { Product } from "@/models/product";
import { Box } from "@mui/material";
import { FC, useState } from "react";
import ProductInput from "./ProductInput";
import ProductDescriptionInput from "./ProductDescriptionInput";
import PriceInput from "./PriceInput";
import DeleteButton from "./DeleteButton";
import IsActiveOption from "./IsActionOption";
import ProductDetails from "./ProductDetails";
import ProductImageUploader from "./ProductImageUploader";

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
    <Box>
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
