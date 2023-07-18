-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 18 Jul 2023 pada 18.14
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
-- Struktur dari tabel `tx_kegiatan_harian`
--

CREATE TABLE `tx_kegiatan_harian` (
  `id` int(11) NOT NULL,
  `tanggal` datetime NOT NULL,
  `jam` datetime NOT NULL,
  `id_data_center` int(11) NOT NULL,
  `id_ruangan` int(11) NOT NULL,
  `kondisi` varchar(50) NOT NULL,
  `id_user_1` int(11) NOT NULL,
  `id_user_2` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tx_kegiatan_harian`
--

INSERT INTO `tx_kegiatan_harian` (`id`, `tanggal`, `jam`, `id_data_center`, `id_ruangan`, `kondisi`, `id_user_1`, `id_user_2`) VALUES
(1, '2023-07-03 07:00:00', '2023-07-03 07:00:00', 3, 12, 'NORMAL', 10, 12),
(3, '2023-07-03 07:00:00', '2023-07-03 07:00:00', 1, 2, 'NORMAL', 10, 28),
(5, '2023-07-03 07:00:00', '2023-07-03 07:00:00', 3, 12, 'ABNORMAL', 10, 12),
(6, '2023-07-03 07:00:00', '2023-07-03 07:00:00', 3, 12, 'ABNORMAL', 10, 12),
(7, '2023-07-10 07:00:00', '2023-07-10 16:23:36', 1, 2, 'NORMAL', 10, 16),
(8, '2023-07-10 07:00:00', '2023-07-10 16:23:54', 1, 9, 'NORMAL', 10, 16);

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
(2, 5, 1, 9, 6, 12, 'adri', 'prosia', '2023-05-29 07:00:00', '2023-05-30 07:00:00', 'cek server mati', 'pelayanan terhambat', 'berhasil up lg', 'SELESAI', 10),
(5, 9, 1, 8, 4, 10, 'wewewewe', 'werwerwer', '2023-06-07 07:00:00', '2023-06-15 07:00:00', 'wetrerterert', 'erterter', 'rtyrtyer', 'PENDING', 23);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `view_kegiatan_dc`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `view_kegiatan_dc` (
`id` int(11)
,`id_kegiatan` int(11)
,`id_data_center` int(11)
,`id_ruangan` int(11)
,`id_item` int(11)
,`id_users` int(11)
,`nama_pic_vendor` varchar(50)
,`nama_perusahaan` varchar(200)
,`tanggal_mulai` varchar(10)
,`tanggal_selesai` varchar(10)
,`deskripsi` text
,`resiko` text
,`hasil` text
,`status` varchar(20)
,`id_user_2` int(11)
,`nama_kegiatan` varchar(250)
,`deskripsi_kegiatan` text
,`nama_dc` varchar(50)
,`lokasi` varchar(50)
,`nama_ruangan` varchar(50)
,`nama_item` varchar(100)
,`deskripsi_item` text
,`nip` varchar(50)
,`nama` varchar(50)
,`no_hp` varchar(20)
,`nip_user2` varchar(50)
,`nama_user2` varchar(50)
,`no_hp_user2` varchar(20)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_dash_kegiatan`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_dash_kegiatan` (
`tahun` int(4)
,`nama_kegiatan` varchar(250)
,`jumlah` bigint(21)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_kondisi_abnormal`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_kondisi_abnormal` (
`tahun` int(4)
,`jumlah` bigint(21)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_kunjungan`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_kunjungan` (
`tahun` int(4)
,`jumlah` bigint(21)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_monitoring_harian`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_monitoring_harian` (
`id` int(11)
,`tahun` int(4)
,`bulan` int(2)
,`tanggal` varchar(10)
,`jam` varchar(10)
,`id_data_center` int(11)
,`id_ruangan` int(11)
,`kondisi` varchar(50)
,`id_user_1` int(11)
,`id_user_2` int(11)
,`nama_dc` varchar(50)
,`lokasi` varchar(50)
,`nama_ruangan` varchar(50)
,`nip` varchar(50)
,`nama` varchar(50)
,`no_hp` varchar(20)
,`nip_user2` varchar(50)
,`nama_user2` varchar(50)
,`no_hp_user2` varchar(20)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_status_pending`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_status_pending` (
`tahun` int(4)
,`jumlah` bigint(21)
);

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `vw_tamu`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `vw_tamu` (
`tahun` int(4)
,`count(nama_perusahaan)` bigint(21)
);

-- --------------------------------------------------------

--
-- Struktur untuk view `view_kegiatan_dc`
--
DROP TABLE IF EXISTS `view_kegiatan_dc`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `view_kegiatan_dc`  AS SELECT `a`.`id` AS `id`, `a`.`id_kegiatan` AS `id_kegiatan`, `a`.`id_data_center` AS `id_data_center`, `a`.`id_ruangan` AS `id_ruangan`, `a`.`id_item` AS `id_item`, `a`.`id_users` AS `id_users`, `a`.`nama_pic_vendor` AS `nama_pic_vendor`, `a`.`nama_perusahaan` AS `nama_perusahaan`, date_format(`a`.`tanggal_mulai`,'%d/%m/%Y') AS `tanggal_mulai`, date_format(`a`.`tanggal_selesai`,'%d/%m/%Y') AS `tanggal_selesai`, `a`.`deskripsi` AS `deskripsi`, `a`.`resiko` AS `resiko`, `a`.`hasil` AS `hasil`, `a`.`status` AS `status`, `a`.`id_user_2` AS `id_user_2`, coalesce(`b`.`nama_kegiatan`,0) AS `nama_kegiatan`, coalesce(`b`.`deskripsi`,0) AS `deskripsi_kegiatan`, `c`.`nama_dc` AS `nama_dc`, `c`.`lokasi` AS `lokasi`, `d`.`nama_ruangan` AS `nama_ruangan`, coalesce(`e`.`nama_item`,0) AS `nama_item`, coalesce(`e`.`deskripsi`,0) AS `deskripsi_item`, `f`.`nip` AS `nip`, `f`.`nama` AS `nama`, `f`.`no_hp` AS `no_hp`, `g`.`nip` AS `nip_user2`, `g`.`nama` AS `nama_user2`, `g`.`no_hp` AS `no_hp_user2` FROM ((((((`tx_kegiatan_piket` `a` left join `ms_kegiatan` `b` on(`b`.`id` = `a`.`id_kegiatan`)) left join `ms_data_center` `c` on(`c`.`id` = `a`.`id_data_center`)) left join `ms_ruangan` `d` on(`d`.`id` = `a`.`id_ruangan`)) left join `ms_item` `e` on(`e`.`id` = `a`.`id_item`)) left join `ms_users` `f` on(`f`.`id` = `a`.`id_users`)) left join `ms_users` `g` on(`g`.`id` = `a`.`id_user_2`))  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_dash_kegiatan`
--
DROP TABLE IF EXISTS `vw_dash_kegiatan`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_dash_kegiatan`  AS SELECT year(`a`.`tanggal_mulai`) AS `tahun`, `b`.`nama_kegiatan` AS `nama_kegiatan`, count(`a`.`id_kegiatan`) AS `jumlah` FROM (`tx_kegiatan_piket` `a` join `ms_kegiatan` `b` on(`b`.`id` = `a`.`id_kegiatan`)) GROUP BY `b`.`nama_kegiatan``nama_kegiatan`  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_kondisi_abnormal`
--
DROP TABLE IF EXISTS `vw_kondisi_abnormal`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_kondisi_abnormal`  AS SELECT year(`tx_kegiatan_harian`.`tanggal`) AS `tahun`, count(`tx_kegiatan_harian`.`kondisi`) AS `jumlah` FROM `tx_kegiatan_harian` WHERE `tx_kegiatan_harian`.`kondisi` = 'ABNORMAL' GROUP BY year(`tx_kegiatan_harian`.`tanggal`)  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_kunjungan`
--
DROP TABLE IF EXISTS `vw_kunjungan`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_kunjungan`  AS SELECT year(`tx_kegiatan_piket`.`tanggal_mulai`) AS `tahun`, count(0) AS `jumlah` FROM `tx_kegiatan_piket` GROUP BY year(`tx_kegiatan_piket`.`tanggal_mulai`)  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_monitoring_harian`
--
DROP TABLE IF EXISTS `vw_monitoring_harian`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_monitoring_harian`  AS SELECT `a`.`id` AS `id`, year(`a`.`tanggal`) AS `tahun`, month(`a`.`tanggal`) AS `bulan`, date_format(`a`.`tanggal`,'%d/%m/%Y') AS `tanggal`, date_format(`a`.`jam`,'%H:%i') AS `jam`, `a`.`id_data_center` AS `id_data_center`, `a`.`id_ruangan` AS `id_ruangan`, `a`.`kondisi` AS `kondisi`, `a`.`id_user_1` AS `id_user_1`, `a`.`id_user_2` AS `id_user_2`, `c`.`nama_dc` AS `nama_dc`, `c`.`lokasi` AS `lokasi`, `d`.`nama_ruangan` AS `nama_ruangan`, `f`.`nip` AS `nip`, `f`.`nama` AS `nama`, `f`.`no_hp` AS `no_hp`, `g`.`nip` AS `nip_user2`, `g`.`nama` AS `nama_user2`, `g`.`no_hp` AS `no_hp_user2` FROM ((((`tx_kegiatan_harian` `a` left join `ms_data_center` `c` on(`c`.`id` = `a`.`id_data_center`)) left join `ms_ruangan` `d` on(`d`.`id` = `a`.`id_ruangan`)) left join `ms_users` `f` on(`f`.`id` = `a`.`id_user_1`)) left join `ms_users` `g` on(`g`.`id` = `a`.`id_user_2`))  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_status_pending`
--
DROP TABLE IF EXISTS `vw_status_pending`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_status_pending`  AS SELECT year(`tx_kegiatan_piket`.`tanggal_mulai`) AS `tahun`, count(`tx_kegiatan_piket`.`status`) AS `jumlah` FROM `tx_kegiatan_piket` WHERE `tx_kegiatan_piket`.`status` = 'PENDING' GROUP BY year(`tx_kegiatan_piket`.`tanggal_mulai`)  ;

-- --------------------------------------------------------

--
-- Struktur untuk view `vw_tamu`
--
DROP TABLE IF EXISTS `vw_tamu`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_tamu`  AS SELECT year(`tx_kegiatan_piket`.`tanggal_mulai`) AS `tahun`, count(`tx_kegiatan_piket`.`nama_perusahaan`) AS `count(nama_perusahaan)` FROM `tx_kegiatan_piket` GROUP BY `tx_kegiatan_piket`.`nama_perusahaan``nama_perusahaan`  ;

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
-- Indeks untuk tabel `tx_kegiatan_harian`
--
ALTER TABLE `tx_kegiatan_harian`
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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

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
-- AUTO_INCREMENT untuk tabel `tx_kegiatan_harian`
--
ALTER TABLE `tx_kegiatan_harian`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `tx_kegiatan_piket`
--
ALTER TABLE `tx_kegiatan_piket`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
