-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 30 Bulan Mei 2023 pada 06.16
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `epiket_v2`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_data_center`
--

CREATE TABLE `ms_data_center` (
  `id` int(11) NOT NULL,
  `nama_dc` varchar(50) NOT NULL,
  `lokasi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_data_center`
--

INSERT INTO `ms_data_center` (`id`, `nama_dc`, `lokasi`) VALUES
(1, 'Data Center Jakarta', 'imigrasi Jakarta'),
(3, 'Disaster Recovery Center', 'Bali');

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_item`
--

CREATE TABLE `ms_item` (
  `id` int(11) NOT NULL,
  `nama_item` varchar(100) NOT NULL,
  `id_ruangan` int(11) NOT NULL,
  `deskripsi` text NOT NULL,
  `parent_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_item`
--

INSERT INTO `ms_item` (`id`, `nama_item`, `id_ruangan`, `deskripsi`, `parent_id`) VALUES
(4, 'Contaiment Bima Sakti', 8, 'kontainer server Bima Sakti', 0),
(5, 'Contaiment Andromeda', 9, 'kontainer server andromeda', 4),
(6, 'Server DPRI', 9, 'server DPRI PUSAT', 5);

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_kegiatan`
--

CREATE TABLE `ms_kegiatan` (
  `id` int(11) NOT NULL,
  `nama_kegiatan` varchar(250) NOT NULL,
  `deskripsi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_kegiatan`
--

INSERT INTO `ms_kegiatan` (`id`, `nama_kegiatan`, `deskripsi`) VALUES
(2, 'Instalasi Server', 'Instalasi server'),
(5, 'Cek Server', 'Cek fisik Server'),
(6, 'Cek Switch', 'Cek fisik Switch'),
(7, 'Rapat', 'Rapat'),
(8, 'Kunjungan', 'Kunjungan Pejabat/Tamu'),
(9, 'Maintenance', 'Maintenance Ruang'),
(17, 'piket lebaran xxxx', 'kondusif aman terkendali'),
(42, 'testing kegiatan', 'tes tes tes ');

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_ruangan`
--

CREATE TABLE `ms_ruangan` (
  `id` int(11) NOT NULL,
  `nama_ruangan` varchar(50) NOT NULL,
  `id_dc` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_ruangan`
--

INSERT INTO `ms_ruangan` (`id`, `nama_ruangan`, `id_dc`) VALUES
(1, 'Ruang Rapat', 1),
(2, 'NOC', 1),
(7, 'CA-KMS', 1),
(8, 'Staging', 1),
(9, 'Ruang Data Center', 1),
(10, 'Ruang UPS', 1),
(11, 'Contaiment Andromeda', 3),
(12, 'Contaiment Andromeda', 3);

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_struktur`
--

CREATE TABLE `ms_struktur` (
  `id` int(11) NOT NULL,
  `nama_struktur` varchar(100) DEFAULT NULL,
  `nip` varchar(50) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_struktur`
--

INSERT INTO `ms_struktur` (`id`, `nama_struktur`, `nip`, `parent_id`, `created_at`, `updated_at`) VALUES
(8, 'Direktorat Sistem dan Teknologi Informasi Keimigrasian', 'xxxxxxxxxxxxxx', 0, NULL, NULL),
(9, 'Sub Direktorat Perencanaan dan Pengembangan', 'xxxxxxxxxxxxxx', 8, NULL, NULL),
(10, 'Sub Direktorat Pemeliharaan dan Pengamanan', 'xxxxxxxxxxxxxx', 8, NULL, NULL),
(11, 'Sub Direktorat Kerjasama dan Pemanfaatan Teknologi Informasi Keimigrasian', 'xxxxxxxxxxxxxx', 8, NULL, NULL),
(12, 'Sub Direktorat Pengelolaan Data dan Pelaporan', 'xxxxxxxxxxxxxx', 8, NULL, NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_users`
--

CREATE TABLE `ms_users` (
  `id` int(11) NOT NULL,
  `nip` varchar(50) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `no_hp` varchar(20) NOT NULL,
  `password` text NOT NULL,
  `id_struktur` int(11) NOT NULL,
  `aktif` int(11) NOT NULL,
  `id_role` int(11) NOT NULL,
  `token` text NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `ms_users`
--

INSERT INTO `ms_users` (`id`, `nip`, `nama`, `no_hp`, `password`, `id_struktur`, `aktif`, `id_role`, `token`, `created_at`, `updated_at`) VALUES
(10, '199404022019012002', 'hana A', '081283628066', '$2a$05$f675x7R7zjJInG9BekNDnOPRnkRn.VqV.a/mEJ9unaMmEQRIMnica', 9, 1, 1, '', '0000-00-00', '0000-00-00'),
(12, '199507252019011001', 'Ruktian Ibnu Wijonarko', '081319886308', '$2a$05$f675x7R7zjJInG9BekNDnOPRnkRn.VqV.a/mEJ9unaMmEQRIMnica', 9, 1, 1, '', '0000-00-00', '0000-00-00'),
(16, '199507252019011002', 'Ruktian Ibnu', '081319886308', '$2a$05$zC1yC1ddLuqxIHS1UZF8SuYRnDWkehvV5uYqfUV3lj0NKMYzmnqvK', 12, 0, 3, '', '0000-00-00', '0000-00-00'),
(23, '199404022019012005', 'testing', '081283628066', '$2a$05$.Ly2RlyY93RUDqqsODxp/uHwgCwJj96pOxDktiGibF5/XjW9Sqwb2', 10, 1, 1, '', '0000-00-00', '0000-00-00'),
(28, '199404022019012008', 'ini testing', '081283628066', '$2a$05$Kk4.QTRXCRhJ1s4KmhWUTu1fthUCWjQZqmjri33lYAxv0jUA73wiW', 10, 1, 1, '', '0000-00-00', '0000-00-00'),
(29, '198109092009011005', 'hidayat', '085330335963', '$2a$05$TIaLubs0mlwduFVHiMDkO.qoSxyDoFfjCF0prorIRRzD14vatOXRq', 9, 1, 1, '', '0000-00-00', '0000-00-00'),
(30, '231564485983', '', '1988012365456', '$2a$05$p1o5lS1GnwkHOtHmo4Htj.500HQZgmxiP3W0OyQ4GW0sul.JDhyU.', 0, 1, 0, '', '0000-00-00', '0000-00-00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tx_kegiatan_piket`
--

CREATE TABLE `tx_kegiatan_piket` (
  `id` int(11) NOT NULL,
  `id_kegiatan` int(11) NOT NULL,
  `id_data_center` int(11) NOT NULL,
  `id_ruangan` int(11) NOT NULL,
  `id_item` int(11) NOT NULL,
  `id_users` int(11) NOT NULL,
  `nama_pic_vendor` varchar(50) NOT NULL,
  `nama_perusahaan` varchar(200) NOT NULL,
  `tanggal_mulai` datetime NOT NULL,
  `tanggal_selesai` datetime NOT NULL,
  `deskripsi` text NOT NULL,
  `resiko` text NOT NULL,
  `hasil` text NOT NULL,
  `status` varchar(20) NOT NULL,
  `id_user_2` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tx_kegiatan_piket`
--

INSERT INTO `tx_kegiatan_piket` (`id`, `id_kegiatan`, `id_data_center`, `id_ruangan`, `id_item`, `id_users`, `nama_pic_vendor`, `nama_perusahaan`, `tanggal_mulai`, `tanggal_selesai`, `deskripsi`, `resiko`, `hasil`, `status`, `id_user_2`) VALUES
(1, 2, 1, 9, 4, 10, 'Arha', 'Syncro', '2023-05-25 00:00:00', '2023-05-25 00:00:00', 'install aplikasi molina', 'gagal install', 'berhasil install', 'SELESAI', 12),
(2, 5, 1, 9, 6, 12, 'adri', 'prosia', '2023-05-29 07:00:00', '2023-05-30 07:00:00', 'cek server mati', 'pelayanan terhambat', 'berhasil up lg', 'SELESAI', 10);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `ms_data_center`
--
ALTER TABLE `ms_data_center`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_item`
--
ALTER TABLE `ms_item`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_kegiatan`
--
ALTER TABLE `ms_kegiatan`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_ruangan`
--
ALTER TABLE `ms_ruangan`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_struktur`
--
ALTER TABLE `ms_struktur`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_users`
--
ALTER TABLE `ms_users`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tx_kegiatan_piket`
--
ALTER TABLE `tx_kegiatan_piket`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `ms_data_center`
--
ALTER TABLE `ms_data_center`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `ms_item`
--
ALTER TABLE `ms_item`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `ms_kegiatan`
--
ALTER TABLE `ms_kegiatan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=43;

--
-- AUTO_INCREMENT untuk tabel `ms_ruangan`
--
ALTER TABLE `ms_ruangan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT untuk tabel `ms_struktur`
--
ALTER TABLE `ms_struktur`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT untuk tabel `ms_users`
--
ALTER TABLE `ms_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT untuk tabel `tx_kegiatan_piket`
--
ALTER TABLE `tx_kegiatan_piket`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
