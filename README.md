# Expense Tracker in Go 🧾

A simple **Expense Tracker** project built using **Golang**, showcasing the use of **pointers**, **functions**, and **slices**. This project helps you understand how these concepts work together while managing a dynamic list of expenses.

---

## 📝 **Features**
- **Add Expenses**: Dynamically add new expenses to the list.
- **Update Expenses**: Modify an existing expense by its index.
- **Calculate Total**: Compute the total of all expenses in the list.

---

## 📂 **Project Structure**
```plaintext
📁 Expense-Tracker
    └── main.go  # Core logic for managing expenses
```

---

## 🛠️ **Tech Used**
- **Go (Golang)**

---

## 🚀 **Getting Started**

### Prerequisites
- [Install Go](https://golang.org/dl/) on your machine.

### Steps to Run
1. Clone the repository or copy the code into a file named `main.go`.
2. Open your terminal and navigate to the project directory.
3. Run the program:
   ```bash
   go run main.go
   ```

---

## 📚 **Concepts Covered**

### **Pointers**
- Used to modify the original `expenses` slice in the functions `addExpense` and `updateExpense`.
- Demonstrates the power of working with references in Go.

### **Functions**
- Modularize operations like adding, updating, and calculating expenses.

### **Slices**
- Dynamically store and manage a list of expenses.
- Use `append` to grow the slice and indexing to access or update elements.

---

## 🕹️ **How It Works**
1. The program starts with an empty list of expenses (`expenses` slice).
2. Use the `addExpense` function to add new expenses.
3. Calculate the total using the `calculateTotal` function.
4. Update specific expenses using the `updateExpense` function.
5. The program demonstrates how changes are reflected dynamically.

---

## 📝 **Example Output**
```plaintext
Added expense: 50.75
Added expense: 25.40
Added expense: 100.00
Current expenses: [50.75 25.4 100]
Total expenses: 176.15
Updated expense at index 1 to 30.00
Updated expenses: [50.75 30 100]
Updated total expenses: 180.75
```

---

## ✨ **Future Enhancements**
- Add input validation to handle incorrect user inputs.
- Allow users to remove expenses from the list.
- Implement a user-friendly CLI menu for better interaction.
- Store expenses in a file for persistent tracking.

---

## 📜 **License**
This project is licensed under the MIT License. Feel free to use, modify, and distribute it.

---

## 💡 **Contributions**
Contributions are welcome! If you'd like to enhance the project, feel free to fork the repository and submit a pull request.

---

Enjoy learning Go with this project! 🌟
