import ProductForm from "@/components/ProductForm/ProductForm";
import Image from "next/image";
import { ChangeEvent } from "react";
import { Product } from "@/models/product";

export default function Home() {
  const mockProduct: Product = {
    name: "Sample Product",
    price: 10.99,
    description: "This is a sample product description",
    id: 0,
    imageUrl: "",
    isActive: false,
  };

  return (
    <main>
      <h1>MyMeal</h1>
      <Image src="/logo.png" alt="MyMeal logo" width={200} height={200} />
      <ProductForm product={mockProduct} />
    </main>
  );
}
