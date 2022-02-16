import { Router } from 'express';
import RoutesExample from './RoutesExample'; 

const router = Router();

router.use('/example', RoutesExample);

export default router;