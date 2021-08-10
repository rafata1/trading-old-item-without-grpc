-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql-40280-db.mysql-40280:280
-- Generation Time: Jul 24, 2021 at 01:14 PM
-- Server version: 8.0.25
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `trading_web_data`
--

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int NOT NULL,
  `owner_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `brand` varchar(255) NOT NULL,
  `type` varchar(255) NOT NULL,
  `amount` int NOT NULL,
  `description` longtext,
  `image_url` mediumtext,
  `status` varchar(45) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `owner_id`, `name`, `brand`, `type`, `amount`, `description`, `image_url`, `status`) VALUES
(1, 1, 'maý sấy tóc', 'panasonic', 'Đồ điện tử', 1, 'Đồ con mới, mới sử dụng tử năm 2019', 'data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAoHCBISEhIREhEYEREREhEYEREREhIYERgZGRUZGhgUGBgcIS4lHB4rHxgaJjgnLDAxODU1GiQ7QDszPy41NTEBDAwMEA8PHhERHD8mIyw1OjQxNTU0Pz8xPT8xPz0/Pzo1NDo/MzcxNDFAMTQ9PzY7P0A/NDQ9NDQxPT81PjFANP/AABEIAQYAwAMBIgACEQEDEQH/xAAcAAEAAgIDAQAAAAAAAAAAAAAAAQcCBgQFCAP/xABKEAABAwICBQgECwQIBwAAAAABAAIDBBEFEgchMUFRBhMiMmFxgZEUQqGxIzNScpKissHC0dIXYnOCFSQlQ1NUg7Q0RGNkhJPh/8QAGQEBAAMBAQAAAAAAAAAAAAAAAAECAwQF/8QAHxEBAAICAgIDAAAAAAAAAAAAAAECAxEEMVFhBSEy/9oADAMBAAIRAxEAPwC5kREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEWNxx9qXHH2oMkWOYcUzDigyRY5hxS4QZIoul0Eooul0Eooul0Eooul0Eooul0EoououEGSLHMOKZhxQZIscw4pmHFBkixuOPtS44+1B8URTY8EEIpseCWPBARLHgpynggIljwWLpGt6zg35xA96D6Iusqcdo4vjKuCP588Tfe5dVV6QMJi61fG7sjzyfYaUGz3S6rXEdMNCzMIIJqjLvs2OM/SJdb+VavX6aat3xFJDF2yOfIfC2QewoLyusJZGsBc9wY0bXOIDfMrzViGknF57g1jo2n1YGsjt3OaM3tWs1ldLM7NNK+Z3ypXuc7zcSg9QVHLHDIyc+IU9xqIbM1xHZZpK4o0hYRs9Pj+jJb7K8wLkU9DLJ8XE+S23m2Od7gg9Q0/LHDZLBuIU9zawdMxhN9gAcRrXdRStcMzXBzflMIcPMLyNPQTMF3wyMHF8b2j2hYUtXJE7PFK+J3yo3ua7zBug9fIvNOF6ScVp7AVTpmj1KkCQHvcel9Zb3gmmeM2bW0xYdQMtMczb7yWO1gdznILaRdZgvKCkrW5qWoZNYAua02kaD8phs5viF2aAiKCg0vShyhfR0bWQkieqfkYQbOa0C73A7tVhfdm3KmcO5N1WIvd6MGOlYzNJ0w0G5sHAnaTex7lvWmasBljhsC5rYsp3tLnPc6x3Xa1oXx0azhtXLbqtpIW3FtrRGXAdznP8boNWk0a4w3ZTh2v1amDz1vCwHILGm7KZ47qmD7pFfHpg4p6YOKChTyGxr/LSf+6P9ag8iMaO2llP+rH+pX36YOKemDig8/nkHi7ttG8974v1J+z7Fv8AJO8Xw/qXoD00cU9NHFBQTdHOLn/kz4zU4971yYdF2LO2wMZ8+eH8Lir09NHFPTBxQUzDokxK4Jkp4z2yvJH0WFdo7RDNlDn1MbngjMyFhBcN9i4hubhfKDvIVpemjinpo4oNEwrRdhjic0tRI9nxkT3Mjc35zA3MBq1EGx3ErY6PR/hUWttGx54yukf7HOI9i7KpdHJbOOk3qSNcWvb8141j7964xr54toFUwb25WVAHa3Ux/hl7ig7Ckwakh+KpYYv4cMbT4kDWueDu3cF0tNyggkOUPyP3xyAskH8rrE+C5npg4oOdmPFcGtwmlnFpqaKb+JFG894JGoqfTBxUirHFBq+J6McLmuWxOp3H1oJHAfRfmaB3ALS8Y0Pzsu6kqWTDXaOUc2+3AOF2k9+VW6KscVPpQ4oPM9fhVbh8jTLHLSyNPQkGZuu3qSNNj4Fb/wAktLU0QEeINNRGLD0hgaJ2j95uoPGzXqPzlaVfWQNjeags5kizxIAYyD6paetfhvVeYjo8p6xsk0AOHmQgwwlpLHAX6cjCbsLidTWnogC4uSAFn4PjFNWR87SzNmZquWnpNNr5XtOtp7CAucvMstDiODVDX3dA+9mTxnNC/svazhq1scO8K9+RHKVuJUolIa2aN2SoYzNkD7Ahzb68rgQdezWNdroKt0m1B/paYEgNDaVvS2C8V7jtuR5X3LUpZQGudGXxvdIA4B2zK1xGUtsbHMDYrc9LkUba8OJ6bxA8tAPUDcjnXtb1CNq02paC6QjqukkLfm5ujfttq8EHybidU3q1U7e6eUfiX3Zygrm7Kubxkc77V1xjGoLEHMPKGvOv0uW44PsPIL70/K7EIz/xJf2SMY4e6/tXV5FjzaDvIeXOIMJLpGSX9V8bbDuyZSvr+0KtOxsPgyT9a10sWAjQbQ3SFWb44T3MkH419W6RanfAw9znj81qRYsTGg3IaR5d9M3wlcPwrP8AaQ/fTDwnP6FpRjUGNBvLNJB30zh3Sg/hC+g0js3wPHc5p+9aAWKCxBv8/L2mkGWSne8cHMjdbuu5fKLldSD4uSpg/dHSZ9ElwWiZFBYgsVnLdg2Veb+JTPHtYuRHy4af76M+ErPtBang+HwPYDYSP2ua5xu08MoOztN/uXZegxbDAz6AWNs8VnUw9PD8Zky0i9bRqfbY2csmn+8i8aho94X2byjkf1Z4WA72OD3eGsBafNgdO8aoyw8Wud7jcLVaymySOYHB4abZhsKtTJFunPyeFk4+pvrU+JW/T1EYeJHudPIOrJI4Oy/MaOizwF12QxrtVDhpGsEg8Qvs2rmb1ZXjue4fetHIvGoxKORjo5Gtex4s5jwHNPeCsuQGHQ01VNzBLY54+lESXC7HXYQ4m+oOeLG+3bqVKMxiqbsqH+Ls3vW+aJsXmmxGJj3lwbFOX3DbHo2bsF9pCDddJnJKSuEE9Pfn4MzXBpAeWE3BbcjWDff6yqGv5O4nG4/1Gpawam2hkc2w3lzQRc7Tr3r00pQeTagVMfxjHx/xGFv2gvkK9/EHvaF63uuNNQQydeGN99ueON3vCDymMQdva0+B/NZjEOLfI/8Axel5eSWGvvmw+m17S2njafNoC6yo0bYO83NGGHiyWdo8g+3sQef21zN7XDyKzbUxn1rd4KuSt0PYc+5jknhO4B7Ht8nNv7VrtdoWnFzBWxv4CaN8ftaXe5BoLS12xwPcQpMfYu2xLRpi0Fz6LzzR61O9kl+5l8/1VrNTDPA7JIx8Lh6kjXtP0XIOcWKMi4cddJssHeGv2LmNrG26bch4XufLaEEGNBETsF+5feOVjuq8E8DqPkV9XRnwQcT0Y77DvcFiYP3m+38lyjEsTGg44gHywPpfkuQ2WYam1Btw5x9vaoMaxLE1tatrV/M6RKZnCxe943jnCR5XXEdHbURbsXKLFlc7DrHA6wmkWtNvuZ24JYoyLmujB2auz8l8XlrdpA7N/kiHGczerQ0HUAdPU1QaWtihbFckkOfI4OJHcGfWVdUdPJUSMigiMr5HhrGAdY7bbeGs8BcnUvR3Izk+3D6RlP0TISZJ3NGoyOAzZf3QAGjsaEHeKVCIJRQpQEREBSoRBN1x66iiqGGOeJkzDtbIxr299jsPavuiDQ8V0VYdNcxc5SOJJ+CeXM18WPvq7AQtCxfRBXxXNPJHVNG4Hm5fouOX6yvlEHk7FMJqKV2SogkgdcgCRjmg2+STqcO0LjRVL2dVxHZfV5L1tUQMkaY5GNex3WY9rXMPe06iqn0ocgaeKnfX0cYhMZaZ4WX5sscQ3OxvqkEi4Gq19ltYVXHirxtDXeFj7FyGYsz1oyO4g/kuoAWRAQdx/ScR3OH8o/NYuxGLg4+A/NdUGX2LdeSGjeqxFgnL201O64ZI9pc99jYlrBa4vcXJGzVdBrT8SZuYT3kD818H4g47GgeZKuSi0M0jbc9VTSkWvzbY42nzDzbxWz4dyAwqn1to2SO1dKfNL9V5LR4BB5+wvDKytfkp4nzOvY5G2Y35ztTWjtJCsLAdDcrrPrqgRN/wqezn7N73DK09wcrlYwNAa0BrRsa0ANHcApQdTgXJujoGZaaBsZIs+TrSu+c86yOzZ2LtkUIClQpQEREBSoRBKKEQSiJdAREQF03LGLnMOrmbSaSot3hjiPaF3K6PlhicdPSS57udMx8UMbGl0kkkjC1rGNGsnX5IPPfJ2lZ/SNLHIxr43VEbXse0OY4OI1Fp1Ea1jyzw5tLX1VOwWYyUmNvyWuAe1o7AHW8F2GFxubidIXANcaqmzBrmuAIe1rm5mki4II2ppOH9r1f/AI/+3jQaqzYe4r1ph8UccMTIwBGyONrANmUMAbbwsvJsR6QXovRni3pOHxscbyUh5l54hrQY3fQc0HtBQbciKEEqERAREQRdTmWKIF0uiIF0uiWQLpdLKbIIussyiyWQTdLoiAtB0s0TnQRTMcWlhkicW2uGyNBvfaAcmU22h1lvy6DlzRmbD6lresxrXi5+Q8PI8WtI8UFEcmG/1zDx/wBxF7JyvrpOP9rVf+h/t41yOSsQdilCxpzATl2oEb3vtY7wFw9I7s2K1x3NlDfosa37kGtMOtXfoXf8FWN/6kDvNjh+BUezaFdWhW9q7hak8/hr/cgtFERARRdQgm6hEQEREBAgClAspREBERAREQEREBa9yxr2Mp5KcdKWoie1rQQA1pFnSOcdTWi+07di2Eqv9IcRD2kEhr423G4ljnbe4OuqZLTWszDq4eKuXNFLdNH5A0pOOU7SQ/I6d7nNByn4F4uOy7gtU5WzZ8Qr33uHVdSR3c663sst60VtzYzK89WKlldfcOlGPvKrOpnMjnyO6z3uc7vc4k+9WjpjkiIyWivW5fKLrBXjoYjtBWP4yxN+jHf8ao+n6wV/6I6fJh7n/wCNVTO8Ghkf4CpZt4REQEREBEUIClYqUGQKlYKUGaLC6XQZosLpdBmiwul0GagrG6EoBK0zSQ6OOKGaR2Voe9hsCXdNpdcAcMntW5qt9MtRlp4GfKdK7yDGD7ZUTETGpa4ss4rxevcNa0c1GSLHKsdaKhOQni5shaPNgVZnYrDwM8zydxSUanVFVDCDxAMZcPJz1XjlLOZ3O30putc7BrPgvTXIygNPh1JE4ZXCFj3jeHvu9wPbmeV555J4aaqsp6e1xLKxr/mA5n/Ua5eoPZ2IhKKEQEREBFCIF0WXNlObKDFFlzZU82UGClZc2U5soMbosubKc2UGKLLmynNlBiiy5spzZQYqodNVV8LBH8mMOP8AO9/6Arg5sqh9MVRfEXMv8XHGLd7A4faKD5Yy7muTuHRdV1TVzzEbyGZ2g91nMPktCct40hnJBg9KdToKBj3jg6TKD7WFaS1qCzdCmFZ6maqcOjTxhjCR68hNyDxDGuH84V0rVtGuBmkw2EOGWSo+GkBuCC8DI032EMDBbjdbXzZQYqFnzZTmygwRZ82VHNlBiiy5spkKD6oiIJRQiCUUIglERARLogIiIC86aUml+M1LBtcaZo8YWW969FqodNVLHz9FKGBsj2T55G5g4hjo8tyAR0cx2jeg07SlPmxJ7BsgigjHDUwO/GuByKwT02up6ci8bn5ptuqNnSf3XAy97guPicrqqZ9RI/M+RxL3NyW1NA1ZdQ1AK3tD2E0jYH1cTnvqXjmp2vyfBkEOLWADqu6Lrm9wBwKCyLeHYNiIoQSihEEqERAREQEREBERAREQEREBERAREQFU2mlwE1DsvzdRbW0HrRga7tP1grZVSaY5L1NK0HW2nlcQC69nPABs117dH5JGrfuDQMDpXT1lJCLnnKiBpF33y5gXHW75IJ27lbJiGG460s6FNjDHBzRYNE7SSCP5j5zFaNoqo+cxWN9tUEc8h1ar5cg9XjJfds87B0sMLKSGqb8ZSVUUjT3XNr7uk1nkg3lEvfWNh2IgIiICIiAiIgIiICIiAiIgIiICIiAiIgKn9LV3VrdpaykjsLtIzOkkJIBGrUBrurgVS6SqSodVSPFNI5hZE2N7I3va4Blz1AcvSc4WNtl96Dp+QFfJQPkrPRHz0rwIppIelJGb582X1gbC9yN2vcdoxzFH45zdDRwStpTIx1XVzRljGsF7sYDtdrvxuBqtcjvtGlGYsPZmY6N75ZHvbIxzX7coOVwBsWsC2wIFkREBERAREQEREBERAREQEREBERAREQEREBSiIIREQEREBERAREQEREH/2Q==', 'available'),
(2, 2, 'PS5', 'Sony', 'Đồ điện tử', 2, 'Đồ con mới, mới sử dụng tử năm 2019', 'https://anphat.com.vn/media/product/33475_sony_playstation_5.jpg', 'available'),
(3, 2, 'PS4', 'Sony', 'Đồ điện tử', 100, 'Đồ con mới, mới sử dụng tử năm 2021', 'pana.jpg', 'available'),
(10, 5, 'Laptop', 'Acer', 'Máy tính', 1, 'Laptop đã cũ, mua từ 2010', 'laptop.jpg', 'available'),
(11, 5, 'Laptop', 'Toshiba', 'Máy tính', 1, 'Laptop đã cũ, cụ nội để lại', 'http://1.bp.blogspot.com/-KVhvcESIKcM/UbBQw8AWdRI/AAAAAAAAADc/4S5tkIuTOis/s1600/1.png', 'available');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int NOT NULL,
  `from_post_id` int NOT NULL,
  `to_post_id` int NOT NULL,
  `status` varchar(45) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `from_post_id`, `to_post_id`, `status`) VALUES
(1, 1, 2, 'negotiating');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `email` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone_number` varchar(45) NOT NULL,
  `gender` varchar(45) NOT NULL,
  `dob` varchar(45) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `username`, `password`, `phone_number`, `gender`, `dob`) VALUES
(1, 'caodung566@gmail.com', 'caodung566', '$2a$10$JzVdABgLjXfeIB86ZU.wMuxL8osbw7F6DMA2G4zcQ5CMTcJoqFKbm', '0974606413', 'male', '2001-03-28'),
(2, 'caodung2803@gmail.com', 'caodung2803', '$2a$10$FPF5lg3bjofPsFcO/H0TauV/d8hZBS1ccjRCFYXDr8ivLNa7aCa0.', '0974606413', 'male', '2001-03-28'),
(3, 'test01@email.com', 'test01', 'abc123', '097777777777', 'other', '\"1949-07-07\"'),
(4, 'test02@gmail.com', 'Nguyen_van_test02', 'abc1234', '0123456789', 'male', '\"1949-07-07\"');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`),
  ADD UNIQUE KEY `username_UNIQUE` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
