select 
b.BC_ID ,
BC_NAME,
BC_IsActive,
SC_Id,
SC_Name,
SC_IsActive
from broad_category as B INNER JOIN sub_category as S
where b.BC_Id = s.BC_Id