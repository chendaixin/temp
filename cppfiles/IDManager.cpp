#include <set>
#define ID_DEFAULT_MAX 1000000

class IDManager
{
public:
    static IDManager* instance()
    {
        if(NULL != s_IDManager)
        {
            return s_IDManager;
        }

        s_IDManager = new IDManager();
        return s_IDManager;
    }

    ~IDManager()
    {
    }

    int getID()
    {
        int iRet = 0;
        if(m_IDset.empty())
        {
            iRet = 1;
            m_IDset.insert(iRet);
            return iRet;
        }

        int iCurrentEnd = *m_IDset.rbegin();
        if(iCurrentEnd < ID_DEFAULT_MAX)
        {
            iRet = iCurrentEnd + 1;
            m_IDset.insert(iRet);
        }
        else
        {
            //查找空闲可用ID
            for(int i=1; i<ID_DEFAULT_MAX; i++)
            {
                if(m_IDset.count(i) == 0)
                {
                    iRet = i;
                    break;
                }
            }
        }

        return iRet;
    }

    int releaseID(int id)
    {
        if(m_IDset.count(id) >0)
        {
            m_IDset.erase(id);
            return 0;
        }

        return -1;
    }

private:
    IDManager()
    {
        printf("IDManager construct\n");
        m_IDset.clear();
        m_IDset.insert(100);
    }

private:
    static IDManager* s_IDManager;
    std::set<int> m_IDset;
};

IDManager* IDManager::s_IDManager = NULL;

int test2()
{
    int id = IDManager::instance()->getID();
    printf("getid:%d\n", id);

    IDManager::instance()->releaseID(id);
    return 0;
}
