import React, { useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import { DefaultApi } from'../../api/apis';
import { Typography,Link } from '@material-ui/core'
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save';
import UndoIcon from '@material-ui/icons/Undo';
import { Alert } from '@material-ui/lab';
import { Link as RouterLink } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TableContainer from '@material-ui/core/TableContainer';
import Paper from '@material-ui/core/Paper';

import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
import { EntGender } from '../../api/models/EntGender';
import { EntPrename } from '../../api/models/EntPrename';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 200,
        },
        selectEmpty: {
          marginTop: theme.spacing(2),
        },
        paper: {
          marginTop: theme.spacing(2),
          marginBottom: theme.spacing(2),
        },
        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
            width: 200,
        },
        table: {
          minWidth: 650,
        },
    }),
  );

export  default  function Create() {
    
  const classes = useStyles();
  const api = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);

  const [patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);
  const [prename, setPrename] = React.useState<EntPrename[]>([]);
  const [gender, setGender] = React.useState<EntGender[]>([]);
  const [medicalrecordstaff, setMedicalrecordstaff] = React.useState<EntMedicalrecordstaff[]>([]);
   
  const [prenameid, setprenameId] = React.useState(Number);
  const [genderid, setgenderId] = React.useState(Number);
  const [medicalrecordstaffid, setmedicalrecordstaffId] = React.useState(Number);
  const [name, setname] = React.useState(String);
  const [idcardnumber, setidcardnumber] = React.useState(String);
  const [age, setage] = React.useState(String);
  const [birthday, setbirthday] = React.useState(String);
  const [bloodtype, setbloodtype] = React.useState(String);
  const [disease, setdisease] = React.useState(String);
  const [allergic, setallergic] = React.useState(String);
  const [phonenumber, setphonenumber] = React.useState(String);
  const [email, setemail] = React.useState(String);
  const [home, sethome] = React.useState(String);
  const [date, setdate] = React.useState(String);
  const [loading, setLoading] = React.useState(true);

  useEffect(() => {
    const getPrename = async () => {
        const res = await api.listPrename({ limit: 5, offset: 2 });
        setLoading(false);
        setPrename(res);
    };
    getPrename();
    const getGender = async () => {
        const res = await api.listGender({ limit: 2, offset: 0 });
        setLoading(false);
        setGender(res);
    };
    getGender();
    const getMedicalrecordstaff = async () => {
        const res = await api.listMedicalrecordstaff({ limit: 1, offset: 0 });
        setLoading(false);
        setMedicalrecordstaff(res);
        console.log(res);
    };
    getMedicalrecordstaff();
    }, [loading]);
    const getPatientrecord = async () => {
        const res = await api.listPatientrecord({ limit: 10, offset: 0 });
        setPatientrecord(res);
      };
    //handle
    const PrenamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setprenameId(event.target.value as number);
      };
      const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setgenderId(event.target.value as number);
      };
      const MedicalrecordstaffhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setmedicalrecordstaffId(event.target.value as number);
      };
      const NamehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setname(event.target.value as string);
      };
      const IdcardnumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setidcardnumber(event.target.value as string);
      };
      const AgehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setage(event.target.value as string);
      };
      const BirthdayhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setbirthday(event.target.value as string);
      };
      const BloodtypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setbloodtype(event.target.value as string);
      };
      const DiseasehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setdisease(event.target.value as string);
      };
      const AllergichandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setallergic(event.target.value as string);
      };
      const PhonenumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setphonenumber(event.target.value as string);
      };
      const EmailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setemail(event.target.value as string);
      };
      const HomehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        sethome(event.target.value as string);
      };
      const DatehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setdate(event.target.value as string);
      };
    
      const CreatePatientrecord = async () => {
        const patientrecord = {
          prename : prenameid,
          gender : genderid ,
          medicalrecordstaff : medicalrecordstaffid,
          name : name,
          idcardnumber : idcardnumber,
          age : age,
          birthday : birthday + ":00+07:00",
          bloodtype : bloodtype,
          disease : disease,
          allergic : allergic,
          phonenumber : phonenumber,
          email : email,
          home : home,
          date : date + ":00+07:00",
        };
        console.log(patientrecord);
        const res: any = await api.createPatientrecord({ patientrecord : patientrecord });
        setStatus(true);
        if (res.id != '') {
          setAlert(true);
        } else {
          setAlert(false);
        }
      };
      const timer = setTimeout(() => {
        setStatus(false);
      }, 1000);
    

    return(
        <Page theme={pageTheme.home}>
            <Header 
            title={`ลงทะเบียนผู้ป่วยนอก`}>
            </Header>
            <Content>
            <br />
            <br />
            <Typography variant="h6" gutterBottom align="center">
              <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>คำนำหน้าชื่อ</InputLabel>
              <Select
                name="prename"
                label="คำนำหน้าชื่อ"
                value={prenameid}
                onChange={PrenamehandleChange}
              >
              {prename.map(item => {
                return (
                  <MenuItem value={item.id}>
                  {item.prefix}
                  </MenuItem>
                  );
                })}
              </Select>
              </FormControl> &emsp;

            <TextField 
            label="ชื่อ-นามสกุล" 
            variant="outlined" 
            value={name}
            onChange={NamehandleChange}
            /> &emsp;

              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                  <Select
                    name="gender"
                    label="เพศ"
                    value={genderid}
                    onChange={GenderhandleChange}
                  >
                  {gender.map(item => {
                    return (
                      <MenuItem value={item.id}>
                      {item.genderstatus}
                      </MenuItem>
                      );
                    })}
                  </Select>
              </FormControl> &emsp;

            <TextField 
            name="idcardnumber"
            label="เลขบัตรประจำตัวประชาชน" 
            variant="outlined" 
            value={idcardnumber}
            onChange={IdcardnumberhandleChange}
            />

            <div className={classes.paper}></div>
            <TextField 
            name="age"
            label="อายุ" 
            variant="outlined" 
            value={age}
            onChange={AgehandleChange}
            /> &emsp;

            <TextField
            name="birthday"
            label="วันเกิด"
            type="datetime-local"
            className={classes.textField}
            value={birthday}
            onChange={BirthdayhandleChange}
            InputLabelProps={{
                shrink: true,
            }}
            /> &emsp; 
            
            <TextField 
            label="กรุ๊ปเลือด" 
            name="bloodtype"
            variant="outlined" 
            value={bloodtype}
            onChange={BloodtypehandleChange}
            /> 
            
            <div className={classes.paper}></div>
            <TextField 
            name="disease"
            label="โรคประจำตัว" 
            variant="outlined" 
            value={disease}
            onChange={DiseasehandleChange}
            /> &emsp;
            
            <TextField
            label="แพ้ยา"  
            name="allergic"
            variant="outlined" 
            value={allergic}
            onChange={AllergichandleChange}
            /> 
            
            <div className={classes.paper}></div>
            <TextField 
            name="phonenumber"
            label="เบอร์โทรที่ติดต่อได้" 
            variant="outlined" 
            value={phonenumber}
            onChange={PhonenumberhandleChange}
            /> &emsp;
            
            <TextField 
            name="email"
            label="อีเมล์" 
            variant="outlined" 
            value={email}
            onChange={EmailhandleChange}
            /> 

            <div className={classes.paper}></div>
            <TextField
            name="home"
            label="ที่อยู่"
            multiline
            variant="outlined"
            style={{ width: "74ch"}}
            value={home}
            onChange={HomehandleChange}
            />
            
            <div className={classes.paper}></div>
            <TextField
            className={classes.formControl}
            name="date"
            label="วันเวลาที่ลงทะเบียนผู้ป่วยนอก"
            type="datetime-local"
            value={date}
            onChange={DatehandleChange}
            InputLabelProps={{
             shrink: true,
             }}
            />&emsp;

              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>พนักงานเวชระเบียน</InputLabel>
              <Select
                label="พนักงานเวชระเบียน"
                value={medicalrecordstaffid}
                onChange={MedicalrecordstaffhandleChange}
              >
              {medicalrecordstaff.map(item => {
                return (
                  <MenuItem value={item.id}>
                  {item.name}
                  </MenuItem>
                  );
                })}
              </Select>
              </FormControl>
              <br />
              <br />
              <FormControl>
              <Button
                onClick={() => {
                  CreatePatientrecord();
                }}
                variant="contained"
                color="primary"
                style={{backgroundColor: "#26c6da"}}
                size="large"
                startIcon={<SaveIcon />}
                >
                บันทึก
              </Button>
              {status ? (
                <div>
                {alert ? (
                  <Alert severity="success">
                    บันทึกสำเร็จ !
                  </Alert>
                  ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
                </div>
              ) : null}
              <br />
              <Link component={RouterLink} to="/Room">
              <Button variant="contained" color="primary" style={{backgroundColor: "#26c6da"}} size="large" startIcon={<UndoIcon />}>
                ย้อนกลับ
              </Button>
              </Link>
              </FormControl>
              <br />
              <br />
              <TableContainer component={Paper}>
              <Table className={classes.table} aria-label="ห้องพัก">
       <TableHead>
         <TableRow>
           <TableCell align="center">หมายเลขผู้ป่วย</TableCell>
           <TableCell align="center">คำนำหน้าชื่อ</TableCell>
           <TableCell align="center">ชื่อ-นามสกุล</TableCell>
           <TableCell align="center">เพศ</TableCell>
           <TableCell align="center">เลขบัตรประจำตัวประชาชน</TableCell>
           <TableCell align="center">อายุ</TableCell>
           <TableCell align="center">วันเกิด</TableCell>
           <TableCell align="center">กรุ๊ปเลือด</TableCell>
           <TableCell align="center">โรคประจำตัว</TableCell>
           <TableCell align="center">แพ้ยา</TableCell>
           <TableCell align="center">เบอร์โทรที่ติดต่อได้</TableCell>
           <TableCell align="center">อีเมล์</TableCell>
           <TableCell align="center">ที่อยู่</TableCell>
           <TableCell align="center">วันเวลาที่ลงทะเบียนผู้ป่วยนอก</TableCell>
           <TableCell align="center">พนักงานเวชระเบียนที่ทำการลงบันทึก</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {patientrecord.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.prename?.prefix}</TableCell>
             <TableCell align="center">{item.name}</TableCell>
             <TableCell align="center">{item.edges?.gender?.genderstatus}</TableCell>
             <TableCell align="center">{item.idcardnumber}</TableCell>
             <TableCell align="center">{item.age}</TableCell>
             <TableCell align="center">{item.bloodtype}</TableCell>
             <TableCell align="center">{item.disease}</TableCell>
             <TableCell align="center">{item.allergic}</TableCell>
             <TableCell align="center">{item.phonenumber}</TableCell>
             <TableCell align="center">{item.email}</TableCell>
             <TableCell align="center">{item.date}</TableCell>
            <TableCell align = "center">{item.edges?.medicalrecordstaff?.name}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
     </TableContainer>
            </Typography>
          </Content>
      </Page>
  );
}