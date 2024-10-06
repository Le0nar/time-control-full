import { FC,  } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
  
interface Row {
    date: string, 
    employeeFullName: string,
    jobTitle: string,
    email: string,
    activityHours: number,
}
  const rows: Row[] = [
    { 
        date: '02.09.2024', 
        employeeFullName: 'Иванов Иван Иванович',
        jobTitle: 'Заместитель заместителя',
        email: 'ivanovii3@mail.ru',
        activityHours: 8,
    },
    { 
        date: '02.09.2024', 
        employeeFullName: 'Иванин Иван Максимович',
        jobTitle: 'Управляющий',
        email: 'ivanovii34@mail.ru',
        activityHours: 7,
    },
    { 
        date: '02.09.2024', 
        employeeFullName: 'Иваненко Иван Рустамович',
        jobTitle: 'Секретарь',
        email: 'ivanovii93@mail.ru',
        activityHours: 4,
    },
    { 
        date: '03.09.2024', 
        employeeFullName: 'Иванов Иван Иванович',
        jobTitle: 'Заместитель заместителя',
        email: 'ivanovii3@mail.ru',
        activityHours: 7,
    },
    { 
        date: '03.09.2024', 
        employeeFullName: 'Иванин Иван Максимович',
        jobTitle: 'Управляющий',
        email: 'ivanovii34@mail.ru',
        activityHours: 6,
    },
    { 
        date: '03.09.2024', 
        employeeFullName: 'Иваненко Иван Рустамович',
        jobTitle: 'Секретарь',
        email: 'ivanovii93@mail.ru',
        activityHours: 3,
    },
    { 
        date: '04.09.2024', 
        employeeFullName: 'Иванов Иван Иванович',
        jobTitle: 'Заместитель заместителя',
        email: 'ivanovii3@mail.ru',
        activityHours: 8,
    },
    { 
        date: '04.09.2024', 
        employeeFullName: 'Иванин Иван Максимович',
        jobTitle: 'Управляющий',
        email: 'ivanovii34@mail.ru',
        activityHours: 4,
    },
    { 
        date: '04.09.2024', 
        employeeFullName: 'Иваненко Иван Рустамович',
        jobTitle: 'Секретарь',
        email: 'ivanovii93@mail.ru',
        activityHours: 6,
    },
    { 
        date: '05.09.2024', 
        employeeFullName: 'Иванов Иван Иванович',
        jobTitle: 'Заместитель заместителя',
        email: 'ivanovii3@mail.ru',
        activityHours: 5,
    },
    { 
        date: '05.09.2024', 
        employeeFullName: 'Иванин Иван Максимович',
        jobTitle: 'Управляющий',
        email: 'ivanovii34@mail.ru',
        activityHours: 2,
    },
    { 
        date: '05.09.2024', 
        employeeFullName: 'Иваненко Иван Рустамович',
        jobTitle: 'Секретарь',
        email: 'ivanovii93@mail.ru',
        activityHours: 3,
    },
  ];

export const ReportPage: FC = () => {

    return (
        <>
            <h1>Сентябрь 2024</h1>
            <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                <TableRow>
                    <TableCell>Дата</TableCell>
                    <TableCell>Сотрудник</TableCell>
                    <TableCell>Должность</TableCell>
                    <TableCell>Email</TableCell>
                    <TableCell>Время активной работы (ч.)</TableCell>
                </TableRow>
                </TableHead>
                <TableBody>
                {rows.map((row) => (
                    <TableRow key={row.date + row.employeeFullName} >
                    <TableCell>{row.date}</TableCell>
                    <TableCell>{row.employeeFullName}</TableCell>
                    <TableCell>{row.jobTitle}</TableCell>
                    <TableCell>{row.email}</TableCell>
                    <TableCell>{row.activityHours}</TableCell>
                    </TableRow>
                ))}
                </TableBody>
            </Table>
            </TableContainer>
        </>
      );
}
