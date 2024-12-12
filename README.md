# Expense Tracker in Go

A simple **Expense Tracker** program written in Go that allows you to manage your expenses effortlessly. You can add, update, and calculate total expenses using this straightforward and efficient program.

---

## 🛠️ **Features**
1. **Add Expenses**: Add new expenses to your list.
2. **Update Expenses**: Modify an existing expense by specifying its index.
3. **Calculate Total**: Automatically calculate and display the total of all expenses.

---

## 📂 **Project Structure**
```plaintext
📁 Expense-Tracker
    └── main.go  # Core logic for the expense tracker
```

---

## 🕹️ **How to Use**

### Prerequisites
- Go installed on your system. Download it [here](https://golang.org/dl/).

### Steps
1. Clone the repository or copy the code into a file named `main.go`.
2. Open your terminal and navigate to the directory containing `main.go`.
3. Run the program:
   ```bash
   go run main.go
   ```

### Sample Execution
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

## 🧑‍💻 **Functions Overview**

### `addExpense`
- **Purpose**: Adds a new expense to the list.
- **Usage**: `addExpense(&expenses, amount)`
- **Parameters**:
  - `expenses`: Pointer to the slice of expenses.
  - `amount`: The expense amount to be added.

### `calculateTotal`
- **Purpose**: Calculates the total of all expenses.
- **Usage**: `calculateTotal(expenses)`
- **Parameters**:
  - `expenses`: Slice of expense amounts.
- **Returns**: The total of all expenses.

### `updateExpense`
- **Purpose**: Updates an existing expense by index.
- **Usage**: `updateExpense(&expenses, index, newAmount)`
- **Parameters**:
  - `expenses`: Pointer to the slice of expenses.
  - `index`: The position of the expense to update.
  - `newAmount`: The new expense amount.

---

## 🤩 **Key Highlights**
- **Dynamic Expense Management**: Easily modify your expenses with no hardcoded limits.
- **Pointer Usage**: Learn how to use pointers to manipulate slices in Go.
- **User-Friendly Output**: Clear, formatted output for a better user experience.

---

## 📜 **License**
This project is licensed under the MIT License. Feel free to use and modify it.

---

## 🚀 **Future Improvements**
- Add a feature to delete expenses.
- Introduce a category system for organizing expenses.
- Implement a CLI menu for easier navigation.

---

## 🌟 **Contribute**
Contributions are welcome! Fork the repository, make your changes, and submit a pull request.

Happy coding! 🎉
