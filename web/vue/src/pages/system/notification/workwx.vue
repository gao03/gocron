<template>
    <el-container>
        <system-sidebar></system-sidebar>
        <el-main>
            <notification-tab></notification-tab>
            <el-form ref="form" :model="form" :rules="rules" size="medium" label-width="100px" style="width: 800px;">
                <el-form-item label="企业ID" prop="corpId">
                    <el-input v-model="form.corp_id" placeholder="CorpId" :maxlength="100" clearable></el-input>
                </el-form-item>
                <el-form-item label="应用密钥" prop="secret">
                    <el-input v-model="form.secret" placeholder="secret" clearable type="password">
                    </el-input>
                </el-form-item>
                <el-form-item label="应用ID" prop="agentId">
                    <el-input v-model="form.agent_id" placeholder="AgentId" clearable>
                    </el-input>
                </el-form-item>
                <br>
                <el-form-item label="模板" prop="template">
                    <el-input type="textarea" :rows="6" placeholder="" v-model="form.template">
                    </el-input>
                </el-form-item>
                <el-form-item size="large">
                    <el-button type="primary" @click="submit">提交</el-button>
                    <el-button @click="resetForm">重置</el-button>
                </el-form-item>
                <el-button type="primary" @click="dialogVisible = true">新增用户</el-button> <br><br>
                <h3>通知用户</h3>
                <div v-for="item in users" :key="item.userId">
                    <el-tag closable @close="deleteUser(item)">{{ item.userId }}</el-tag>
                </div>
            </el-form>
            <el-dialog title="" :visible.sync="dialogVisible" width="30%">
                <el-form :model="form">
                    <el-form-item label="企微用户ID">
                        <el-input v-model.trim="newUserId"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="saveUser">确 定</el-button>
                    </el-form-item>
                </el-form>
            </el-dialog>
        </el-main>
    </el-container>
</template>
<script>
import systemSidebar from '../sidebar'
import notificationTab from './tab'
import notificationService from '../../../api/notification'
export default {
    name: 'notification-workwx',
    components: { notificationTab, systemSidebar },
    props: [],
    data() {
        return {
            users: [],
            newUserId: "",
            dialogVisible: false,
            form: {
                corp_id: "",
                secret: "",
                agent_id: undefined,
                template: "",
            },
            rules: {
                corp_id: [{
                    required: true,
                    message: 'CorpId',
                    trigger: 'blur'
                }],
                secret: [{
                    required: true,
                    message: 'secret',
                    trigger: 'blur'
                }],
                agent_id: [{
                    required: true,
                    message: 'AgentId',
                    trigger: 'blur'
                }],
                template: [{
                    required: true,
                    message: 'AgentId',
                    trigger: 'blur'
                }],
            },
        }
    },
    computed: {},
    watch: {},
    created() {
        this.init()
    },
    mounted() { },
    methods: {
        submit() {
            this.$refs['form'].validate((valid) => {
                if (!valid) {
                    return false
                }
                this.save()
            })
        },
        save() {
            notificationService.updateWorkwx(this.form, () => {
                this.$message.success('更新成功')
                this.init()
            })
        },
        resetForm() {
            this.$refs['elForm'].resetFields()
        },
        saveUser() {
            if (this.newUserId === '') {
                this.$message.error('参数不完整')
                return
            }
            notificationService.createWorkwxUser({
                userId: this.newUserId
            }, () => {
                this.dialogVisible = false
                this.init()
            })
        },
        deleteUser(item) {
            notificationService.removeWorkwxUser(item.id, () => {
                this.init()
            })
        },
        init() {
            notificationService.workwx((data) => {
                this.form.corp_id = data.corpId
                this.form.secret = data.secret
                this.form.agent_id = data.agentId
                this.form.template = data.template
                this.users = data.users
            })
        }
    }
}

</script>
<style>

</style>
