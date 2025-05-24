CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('INCOME', 'EXPENSE', 'INVESTMENT')),
    description TEXT,
    month_year DATE NOT NULL
);


INSERT INTO transactions (date, amount, type, description, month_year) VALUES
('2025-05-01 10:00:00', 2500.00, 'INCOME', 'Salary for May', '2025-05-01'),
('2025-05-03 14:30:00', 150.75, 'EXPENSE', 'Groceries', '2025-05-01'),
('2025-05-10 09:00:00', 500.00, 'INVESTMENT', 'Stocks purchase', '2025-05-01'),
('2025-04-28 11:15:00', 120.00, 'EXPENSE', 'Internet bill', '2025-04-01'),
('2025-04-15 08:00:00', 2400.00, 'INCOME', 'Freelance project', '2025-04-01');
