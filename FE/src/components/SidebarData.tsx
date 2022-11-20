import React from "react";
import HomeIcon from "@mui/icons-material/Home";
import AttachEmailIcon from "@mui/icons-material/AttachEmail";
import AssessmentIcon from "@mui/icons-material/Assessment";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import AddCardIcon from "@mui/icons-material/AddCard";
import BackupIcon from "@mui/icons-material/Backup";
import SettingsIcon from "@mui/icons-material/Settings";

export const SidebarData = [
  {
    title: "ホーム",
    icon: <HomeIcon />,
    link: "/pointlist/",
  },
  {
    title: "メール",
    icon: <AttachEmailIcon />,
    link: "/pointlist/",
  },
  {
    title: "アナリティクス",
    icon: <AssessmentIcon />,
    link: "/pointlist/",
  },
  {
    title: "友達追加",
    icon: <PersonAddIcon />,
    link: "/pointlist/",
  },
  {
    title: "お支払い設定",
    icon: <AddCardIcon />,
    link: "/pointlist/",
  },
  {
    title: "アップロード",
    icon: <BackupIcon />,
    link: "/pointlist/",
  },
  {
    title: "詳細設定",
    icon: <SettingsIcon />,
    link: "/pointlist/",
  },
];