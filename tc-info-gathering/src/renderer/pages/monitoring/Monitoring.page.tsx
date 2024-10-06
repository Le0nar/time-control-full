import { FC, useState } from "react";
import { FormControl, InputLabel, Select, MenuItem, Button } from "@mui/material";
import './Monitoring.page.css';

const TASK_LIST = [
    'ЗАДАЧА 1',
    'ЗАДАЧА 2',
    'ЗАДАЧА 3',
]

export const MonitoringPage: FC = () => {
    const [task, setTask] = useState('')
    const [isActive, setIsActive] = useState(false)

    return (
        <div className="monitoring-container">
            <FormControl>
                <InputLabel id="demo-simple-select-label">Задача</InputLabel>
                <Select
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={task}
                    label="Задача"
                    onChange={(event) => setTask(event.target.value)}
                    placeholder="Выберите задачу"
                >
                    {TASK_LIST.map((taskName) =>  (
                        <MenuItem value={taskName} key={taskName}>{taskName}</MenuItem>
                    ))}
                </Select>
            </FormControl>
            <Button 
                onClick={() => setIsActive(!isActive)} 
                color={isActive ? 'warning' : 'success'}
                variant="contained"
            >
               {isActive ? 'Выключить мониторинг' : 'Включить мониторинг'}
            </Button>
        </div>
    )
}
